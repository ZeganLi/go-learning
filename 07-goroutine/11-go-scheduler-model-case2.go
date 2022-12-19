package main

import (
	"fmt"
	"runtime"
	"time"
)

// 如果一定要让deadloop goroutine以外的goroutine无法得到调度，该如何做呢？
// 一种思路是：让Go运行时不要启动那么多P，让所有用户级的goroutine都在一个P上被调度。
// 下面是实现上述思路的三种办法：
// 1、在main函数的最开头处调用runtime.GOMAXPROCS(1)；
// 2、设置环境变量export GOMAXPROCS=1后再运行程序；
// 3、找一台单核单线程的机器（不过现在这样的机器太难找了，只能使用云服务器实现）。

// Go 1.14版本中加入了goroutine的抢占式调度，新的调度方式利用操作系统信号机制，因此在Go 1.14及后续版本中，该例子将不适用
func main() {
	runtime.GOMAXPROCS(1)
	go dead_loop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}

func dead_loop() {
	for {

	}
}
