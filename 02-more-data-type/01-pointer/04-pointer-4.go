package main

import "fmt"

/**
值类型，都有其对应的指针类型。形式为 *数据类型,比如int的指针类型就算*int，float32对应的指针类型就算*float32
在golang中，
值类型包括 int系列 float系列 bool string 数组 和结构体
引用类型包括 指针、slice切片、map、管道cha、interface等

值类型，变量存储值，通常在栈上分配。
引用类型，变量存储地址值，这个地址值对应的空间才真正的存储数据。内存通常在堆中分配。当没有任何变量引用这个地址时，由GC回收
*/
func main() {
	// 声明实际变量
	var a int = 20
	// 声明指针变量
	var p *int
	// 给指针变量赋值，将变量a的地址赋值给p
	p = &a

	// 打印a的类型和值
	fmt.Printf("a 的类型是%T,值是%d \n", a, a)

	// 打印&a的类型和值
	fmt.Printf("a 的类型是%T,值是%d \n", &a, &a)

	// 打印p的类型和值
	fmt.Printf("p 的类型是%T,值是%d \n", p, p)

	// 打印&p的类型和值
	fmt.Printf("&p 的类型是%T,值是%d \n", &p, &p)

	// 打印*p的类型和值
	fmt.Printf("*p 的类型是%T,值是%d \n", *p, *p)

	// 打印*&a的类型和值
	fmt.Printf("*&a 的类型是%T,值是%d \n", *&a, *&a)

	// 打印*&a的类型和值
	fmt.Printf("*&p 的类型是%T,值是%d \n", *&p, *&p)

	fmt.Println(a, &a, *&a)
	fmt.Println(p, &p, *p, *(&p), &(*p))
}
