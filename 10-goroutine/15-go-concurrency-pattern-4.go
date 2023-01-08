package main

import (
	"fmt"
	"sync"
	"time"
)

//支持超时机制的等待

//我们通过一个定时器（time.Timer）设置了超时等待时间，并通过select原语同时监听timer.C和done这两个channel，哪个先返回数据就执行哪个case分支
func main() {
	done := spawnGroup(5, worker, 30)
	println("spawn a group of workers")
	timer := time.NewTimer(time.Second * 5)
	defer timer.Stop()
	select {
	case <-timer.C:
		println("wait group workers exit timeout!")
	case <-done:
		println("group workers done")
	}
}

func worker(args ...interface{}) {
	if len(args) == 0 {
		return
	}

	interval, ok := args[0].(int)
	if !ok {
		return
	}

	time.Sleep(time.Second * time.Duration(interval))
}

//通过sync.WaitGroup，spawnGroup每创建一个goroutine都会调用wg.Add(1)，新创建的goroutine会在退出前调用wg.Done。
//在spawnGroup中还创建了一个用于监视的goroutine，该goroutine调用sync.WaitGroup的Wait方法来等待所有goroutine退出。
//在所有新创建的goroutine退出后，Wait方法返回，该监视goroutine会向done这个channel写入一个信号，
//这时main goroutine才会从阻塞在done channel上的状态中恢复，继续往下执行
func spawnGroup(n int, f func(args ...interface{}), args ...interface{}) chan struct{} {
	c := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			name := fmt.Sprintf("worker-%d:", i)
			f(args...)
			println(name, "done")
			wg.Done() // worker done!
		}(i)
	}

	go func() {
		// 关键在这个地方，启动一个goroutine，来等待sync.WaitGroup的完成，等待完成之后再向channel发送数据，来通知main goroutine完成
		wg.Wait()
		c <- struct{}{}
	}()

	return c
}
