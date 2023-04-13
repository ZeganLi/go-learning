package main

import (
	"fmt"
	"time"
)

// select语句的机制有点像switch语句，不同的是，select会随机挑选一个可通信的case来执行，
// 如果所有case都没有数据到达，则执行default，如果没有default语句，select就会阻塞，直到有case接收到数据。
func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() { ch1 <- 100 }()
	go func() { ch1 <- 200 }()

	time.Sleep(1 * time.Second)

	// 输出结果在100和200之间随机选择，可以证明select的选择是随机的
	select {
	case data := <-ch1:
		fmt.Println("ch1读取到了数据:", data)
	case data := <-ch2:
		fmt.Println("ch2读取到了数据:", data)
	default:
		fmt.Println("执行到了默认分支")

	}
}
