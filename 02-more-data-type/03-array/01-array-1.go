package main

import (
	"fmt"
)

/**
类型 [n]T 表示拥有 n 个 T 类型的值的数组。
表达式

var a [10]int

会将变量 a 声明为拥有 10 个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组
*/
func main() {
	// 声明数组
	var a [2]string
	fmt.Println(a) // [ ]

	//fmt.Println(a == nil)    // Cannot convert 'nil' to type '[2]string'  没有空数组
	fmt.Println(len(a) == 0) // false 只有没有元素的数组

	/**
	Go语言中的数组并非引用类型，而是值类型。当它们被分配给一个新变量时，会将原始数组复制出一份分配给新变量。
	因此对新变量进行更改，原始数组不会有反应。
	将数组作为函数参数进行传递，它们将通过值传递，原始数组依然保持不变。因此如果是大数组的函数传递，会非常消耗性能。
	在C和Java中，可以通过传递数组的引用或指针，但GO里面使用slice切片
	*/
	fmt.Printf("数组a的类型为：%T \n", a) // [2]string

	// 为数组赋值，修改数组中的元素，索引从0开始
	a[0] = "hello"
	a[1] = "world"

	//a[2] = "Invalid array index '2' " // 如果超出数组的最大会编译错误
	fmt.Println(a[0], a[1])
	fmt.Println(a) // [hello world] 注意，是没有“,”号的！Java中是有的！

	/**
	初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
	如果忽略 [] 中的数字，不设置数组长度，Go语言会根据元素的个数来设置数组的长度。可以忽略声明中数组的长度并将其替换为“…”
	*/
	its1 := [6]int{1, 2, 3, 4, 5, 6}
	its2 := []int{1, 2, 3, 4, 5, 6}
	its3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(its1) // [1 2 3 4 5 6]
	fmt.Println(its2) // [1 2 3 4 5 6]
	fmt.Println(its3) // [1 2 3 4 5 6]
}
