package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpu := runtime.NumCPU()
	call := runtime.NumCgoCall()
	goroutine := runtime.NumGoroutine()
	fmt.Printf("当前进程可用的CPU数: %d \n", cpu)
	fmt.Printf("当前进程调用Cgo的次数: %d \n", call)
	fmt.Printf("当前存在的goroutine的数量: %d \n", goroutine)

}
