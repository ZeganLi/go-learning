package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(10 * time.Millisecond)
		data := <-ch1
		fmt.Println("ch1:", data)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		data := <-ch2
		fmt.Println("ch2:", data)
	}()

	select {
	case ch1 <- 100:
		close(ch1)
		fmt.Println("ch1写入")

	case ch2 <- 200:
		close(ch2)
		fmt.Println("ch2写入")

	case <-time.After(2 * time.Second):
		fmt.Println("执行延时通道")

		// default:
		// 	fmt.Println("执行到默认通道")

	}

	time.Sleep(4 * time.Second)
	fmt.Println("main over")
}
