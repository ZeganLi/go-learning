package main

import (
	"errors"
	"fmt"
	"time"
)

var OK = errors.New("ok")

//获取goroutine的退出状态
// 其实就是返回了一个error类型的channel，用channel来承载需要通信的数据类型
func main() {

	done := spawn(worker, 5)
	println("spawn worker1")
	err := <-done
	fmt.Println("worker1 done:", err)

	done = spawn(worker)
	println("spawn worker2")
	err = <-done
	fmt.Println("worker2 done:", err)
}

func worker(args ...interface{}) error {
	if len(args) == 0 {
		return errors.New("invalid args")
	}

	interval, ok := args[0].(int)
	if !ok {
		return errors.New("invalid interval arg")
	}

	time.Sleep(time.Second * (time.Duration(interval)))
	return OK
}

//我们将channel中承载的类型由struct{}改为了error，这样channel承载的信息就不只是一个信号了，还携带了有价值的信息：新goroutine的结束状态
func spawn(f func(args ...interface{}) error, args ...interface{}) chan error {
	c := make(chan error)
	go func() { c <- f(args...) }()
	return c
}
