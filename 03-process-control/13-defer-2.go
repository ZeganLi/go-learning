package main

import "fmt"

func main() {
	//推迟的函数调用会被压入一个栈中。当外层函数返回时，会按照后进先出的顺序调用
	fmt.Println("counting.")
	for i := 0; i < 9; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("end.")

	//counting.
	//end.
	//8
	//7
	//...
	//0
}
