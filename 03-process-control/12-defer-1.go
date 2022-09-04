package main

import "fmt"

func main() {
	// defer 语句会将函数推迟到外函数结束的时候执行
	// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	defer fmt.Println("world")

	fmt.Println("hello")
}
