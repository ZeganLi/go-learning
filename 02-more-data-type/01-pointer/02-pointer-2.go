package main

import "fmt"

func main() {

	var x int = 10

	fmt.Printf("x变量的地址是：%d\n", &x) // 824633811112
	var p *int
	p = &x
	fmt.Printf("p指针对应的值是：%d\n", *p)   // 10
	fmt.Printf("p指针对应的地址值是：%d\n", &p) // 824633745456

	// 指针拥有类型，golang中不允许将一个int类型的指针给到其他类型的指针！！！
	//var pf *float64
	//pf = &x

}
