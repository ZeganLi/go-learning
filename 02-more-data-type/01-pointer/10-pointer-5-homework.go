package main

import "fmt"

func main() {
	//获取一个int变量num的地址，并通过指针去修改其值
	homeWork1()
	homeWork2()
}

func homeWork2() {
	var a int = 300 // 300
	var b int = 400 // 400
	var ptr *int = &a
	*ptr = 100 // a=100
	ptr = &b
	*ptr = 200 // b=200

	fmt.Printf("a = %d,b=%d,*ptr=%d", a, b, *ptr) // 100,200,200
}

func homeWork1() {
	var num int = 20

	fmt.Printf("num的值为:%v\n", num)
	fmt.Printf("num的地址值为:%v\n", &num)

	var ptr *int = &num

	*ptr = 100

	fmt.Println("ptr所引用的值: ", *ptr)
	fmt.Printf("num的值为:%v\n", num)
}
