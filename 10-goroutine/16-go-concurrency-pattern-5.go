package main

import "time"

//notify-and-wait模式

//很多时候，goroutine创建者需要主动通知那些新goroutine退出，尤其是当main goroutine作为创建者时。
//main goroutine退出意味着Go程序的终止，而粗暴地直接让maingoroutine退出的方式可能会导致业务数据损坏、不完整或丢失。
//我们可以通过notify-and-wait（通知并等待）模式来满足这一场景的要求。虽然这一模式也不能完全避免损失，但是它给了各个goroutine一个挽救数据的机会，从而尽可能减少损失

//通知并等待一个goroutine退出
func main() {
	quit := spawn(worker)
	println("spawn a worker goroutine")
	time.Sleep(5 * time.Second)
	// 通知新创建的goroutine退出
	println("notify the worker to exit...")
	quit <- "exit"
	timer := time.NewTimer(time.Second * 10)
	defer timer.Stop()
	select {
	case status := <-quit:
		println("worker done:", status)

	case <-timer.C:
		println("wait worker exit timeout")

	}
}

func worker(j int) {
	time.Sleep(time.Second * time.Duration(j))
}

//使用创建模式创建goroutine的spawn函数返回的channel的作用发生了变化，从原先的只是用于新goroutine发送退出信号给创建者，
//变成了一个双向的数据通道：既承载创建者发送给新goroutine的退出信号，也承载新goroutine返回给创建者的退出状态
func spawn(f func(int2 int)) chan string {
	quit := make(chan string)
	go func() {
		var job chan int // 模拟job channel
		for {
			select {
			case j := <-job:
				f(j)
			case <-quit:
				quit <- "ok"

			}

		}
	}()

	return quit
}
