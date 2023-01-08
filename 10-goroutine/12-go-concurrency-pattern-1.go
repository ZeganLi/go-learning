package main

import "time"

// 等待一个goroutine退出
func main() {
	done := spawn(worker, 5)
	println("spawn a worker goroutine")

	// main goroutine在创建完新goroutine后便在该channel上阻塞等待，直到新goroutine退出前向该channel发送了一个信号。
	<-done
	println("worker done")
}

func worker(args ...interface{}) {
	if len(args) == 0 {
		return
	}

	interval, ok := args[0].(int)
	if !ok {
		return
	}

	time.Sleep(time.Second * (time.Duration(interval)))
}

// spawn函数使用典型的goroutine创建模式创建了一个goroutine，main goroutine作为创建者通过spawn函数返回的channel与新goroutine建立联系，
// 这个channel的用途就是在两个goroutine之间建立退出事件的“信号”通信机制
func spawn(f func(args ...interface{}), args ...interface{}) chan struct{} {
	c := make(chan struct{})
	go func() {
		f(args...)
		c <- struct{}{}
	}()

	return c
}
