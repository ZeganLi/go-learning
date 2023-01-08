package main

import (
	"fmt"
	"sync"
	"time"
)

//通知并等待多个goroutine退出

//Go语言的channel有一个特性是，当使用close函数关闭channel时，所有阻塞到该channel上的goroutine都会得到通知
func main() {
	quit := spawnGroup(5, worker)
	println("spawn a group of workers")
	time.Sleep(5 * time.Second)

	// 通知 worker goroutine 组退出
	println("notify the worker group to exit...")
	quit <- struct{}{}

	timer := time.NewTimer(time.Second * 5)

	defer timer.Stop()
	select {
	case <-timer.C:
		println("wait group workers exit timeout!")
	case <-quit:
		println("group workers done")
	}
}

func worker(j int) {
	time.Sleep(time.Second * time.Duration(j))

}

//创建者直接利用了worker goroutine接收任务（job）的channel来广播退出通知，而实现这一广播的代码就是close(job)。此时各个worker goroutine监听jobchannel，当创建者关闭job channel时，通过“comma ok”模式获取的ok值为false
//，也就表明该channel已经被关闭，于是worker goroutine执行退出逻辑（退出前wg.Done()被执行）
func spawnGroup(n int, f func(int)) chan struct{} {

	quit := make(chan struct{})
	job := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done() // 保证wg.Done在goroutine推出前被执行
			name := fmt.Sprintf("worker-%d:", i)
			for {
				j, ok := <-job
				if !ok {
					println(name, "done")
					return
				}

				// 执行这个job
				// 实际不会执行到这里，因为阻塞在<-job的goroutine，接到channel被关闭的消息，ok变为false
				worker(j)
			}
		}(i)

	}

	go func() {
		<-quit
		close(job)
		wg.Wait()
		quit <- struct{}{}
	}()

	return quit
}
