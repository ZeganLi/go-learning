package main

import (
	"fmt"
	"time"
)

// 尽管存在运行着死循环的deadloop goroutine，maingoroutine仍然得到了调度

// 假设deadloop goroutine被调度到P1上，P1在M1（对应一个操作系统线程）上运行；
// 而main goroutine被调度到P2上，P2在M2上运行，M2对应另一个操作系统线程，而线程在操作系统调度层面被调度到物理的CPU核上运行。
// 我们有多个CPU核，因此即便deadloop占满一个核，我们还可以在另一个CPU核上运行P2上的main goroutine，这也是main goroutine得到调度的原因。
func main() {

	go deadloop()

	for {
		time.Sleep(time.Second)
		fmt.Println("I got scheduled!")

	}
}

func deadloop() {
	for {

	}
}
