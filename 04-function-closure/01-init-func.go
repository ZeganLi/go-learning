package main

import "fmt"

/**
go中有两个特殊的函数：一个是main包中的main函数，它是所有Go可执行程序的入口函数；另一个是init函数
*/
func main() {

	//init() 无法显示的调用
}

/**
init函数：无参无返回值，
一个Go包可以拥有多个init函数，
每个组成Go包的Go源文件中可以定义多个init函数。

Go运行时不会并发调用init函数，它会等待一个init函数执行完毕并返回后再执行下一个init函数，
且每个init函数在整个Go程序生命周期内仅会被执行一次。
因此，init函数极其适合做一些包级数据的初始化及初始状态的检查工作。

一般来说，先被传递给Go编译器的源文件中的init函数先被执行，
同一个源文件中的多个init函数按声明顺序依次执行。
但Go语言的惯例告诉我们：不要依赖init函数的执行次序！！！
*/
func init() {
	fmt.Println("inti——1")
}

func init() {
	fmt.Println("inti——2")

}
