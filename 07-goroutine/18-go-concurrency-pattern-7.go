package main

import (
	"errors"
	"sync"
	"time"
)

//一组goroutine的退出 总体上有两种情况。
//一种是并发退出 在这类退出方式下，各个goroutine的退出先后次序对数据处理无影响，因此各个goroutine可以并发执行退出逻辑；
//另一种则是串行退出即各个goroutine之间的退出是按照一定次序逐个进行的，次序若错了可能会导致程序的状态混乱和错误
func main() {

}

//凡是实现了该接口的类型均可在程序退出时得到退出的通知和调用，从而有机会做退出前的最后清理工作
type GracefullyShutdowner interface {
	Shutdown(waitTimeout time.Duration) error
}

type ShutdownerFunc func(time.Duration) error

func (f ShutdownerFunc) Shutdown(waitTimeout time.Duration) error { return f(waitTimeout) }

// 并发退出

// ConcurrentShutdown
//我们将各个GracefullyShutdowner接口的实现以一个变长参数的形式传入ConcurrentShutdown函数。
//ConcurrentShutdown函数的实现也很简单（类似上面的超时等待多个goroutine退出的模式），具体如下：
//1）为每个传入的GracefullyShutdowner接口实现的实例启动一个goroutine来执行退出逻辑，并将timeout参数传入每个实例的Shutdown方法中；
//2）通过sync.WaitGroup在外层等待每个goroutine的退出；
//3）通过select监听一个退出通知channel和一个timer channel，决定到底是正常退出还是超时退出。
func ConcurrentShutdown(waitTimeout time.Duration, shutdowners ...GracefullyShutdowner) error {
	c := make(chan struct{})

	go func() {
		var wg sync.WaitGroup
		for _, g := range shutdowners {
			wg.Add(1)
			go func(shutdowner GracefullyShutdowner) {
				defer wg.Done()
				shutdowner.Shutdown(waitTimeout)
			}(g)
		}
		wg.Wait()
		c <- struct{}{}
	}()

	timer := time.NewTimer(waitTimeout)
	defer timer.Stop()
	select {
	case <-c:
		return nil
	case <-timer.C:
		return errors.New("wait timeout")
	}
}

// 串行退出

//串行退出有个问题是waitTimeout值的确定，因为这个超时时间是所有goroutine的退出时间之和
//在上述代码里，将每次的left（剩余时间）传入下一个要执行的goroutine的Shutdown方法中。
//select同样使用这个left作为timeout的值（通过timer.Reset重新设置timer定时器周期）
func SequentialShutdown(waitTimeout time.Duration, shutdowners ...GracefullyShutdowner) error {
	start := time.Now()
	var left time.Duration
	timer := time.NewTimer(waitTimeout)

	for _, g := range shutdowners {
		elapsed := time.Since(start)
		left = waitTimeout - elapsed
		c := make(chan struct{})
		go func(shutdowner GracefullyShutdowner,
		) {
			shutdowner.Shutdown(left)
			c <- struct{}{}
		}(g)
		timer.Reset(left)
		select {
		case <-c: // 继续执行
		case <-timer.C:
			return errors.New("wait timeout")
		}
	}
	return nil
}
