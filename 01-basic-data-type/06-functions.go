package main

import "fmt"

//当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
func add1(x, y int) int {
	return x + y
}

/*
1、函数可以没有参数或接受多个参数。
2、类型在变量名之后
*/
func add2(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add1(10, 20))
	fmt.Println(add2(10, 20))
}
