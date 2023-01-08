package main

import "fmt"

/**
信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。

ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
（“箭头”就是数据流的方向。）

创建channel
ch := make(chan int)

发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。
*/
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c) // 7,2,8
	go sum(s[len(s)/2:], c) // -9,4,0

	x, y := <-c, <-c
	fmt.Println(x, y, x+y) // -5 17 12
	//println(x, y, x+y) // 17,-5,12
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
