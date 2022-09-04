package main

import (
	"fmt"
	"net"
	"sync"
)

/*
没有明确初始值的变量声明会被赋予它们的 零值。

零值是：

所有整型类型：0
浮点类型：0.0
布尔类型：false
字符串类型：""
指针、interface、切片（slice）、channel、map、function：nil

Go的零值初始是递归的，即数组、结构体等类型的零值初始化就是对其组成元素逐一进行零值初始化
*/

func main() {
	var i int
	var f float64
	var b bool
	var s string

	//var a uint = -1000 // constant -1000 overflows uint
	//fmt.Printf("数据类型为：%T,其字节长度为： %d\n", a, unsafe.Sizeof(a))

	//var c uint8 = 256 // constant 256 overflows uint8
	//fmt.Printf("数据类型为：%T,其字节长度为： %d\n", c, unsafe.Sizeof(c))

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	/**
	Go从诞生以来就一直秉承着尽量保持“零值可用”的理念
	*/
	fmt.Println("零值可用1---------------------------------------------------------")

	// 声明了一个[]int类型的切片zeroSlice，但并没有对其进行显式初始化，这样zeroSlice这个变量就被Go编译器置为零值nil。
	// 按传统的思维，对于值为nil的变量，我们要先为其赋上合理的值后才能使用。但由于Go中的切片类型具备零值可用的特性，
	// 我们可以直接对其进行append操作，而不会出现引用nil的错误
	var zeroSlice []int
	zeroSlice = append(zeroSlice, 1)
	zeroSlice = append(zeroSlice, 2)
	zeroSlice = append(zeroSlice, 3)
	fmt.Println(zeroSlice) // 输出：[1 2 3]

	fmt.Println("零值可用2---------------------------------------------------------")
	var p *net.TCPAddr
	fmt.Println(p) //输出：<nil>

	//不过Go并非所有类型都是零值可用的，并且零值可用也有一定的限制，比如：在append场景下，零值可用的切片类型不能通过下标形式操作数据
	var sa []int
	sa[0] = 12          // 报错！
	sa = append(sa, 12) // 正确

	//另外，像map这样的原生类型也没有提供对零值可用的支持：
	var m map[string]int
	m["go"] = 1 // 报错！

	m1 := make(map[string]int)
	m1["go"] = 1 // 正确

	// 零值可用的类型要注意尽量避免值复制：
	var mu sync.Mutex
	//mu1 := mu // 错误:避免值复制
	foo(mu) // 错误: 避免值复制

	// 可以通过指针方式传递类似Mutex这样的类型
	var muPtr sync.Mutex
	fooPtr(&muPtr) // 正确
}

func foo(mu sync.Mutex) {

}

func fooPtr(mu *sync.Mutex) {

}
