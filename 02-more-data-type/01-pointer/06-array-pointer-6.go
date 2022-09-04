package main

import "fmt"

// 指针数组就是元素为指针类型的数组
// 指针数组的定义
// var ptr [3]*string
// 变量ptr即为一个大小为3个元素的string类型指针的数组

const COUNT = 4

func main() {

	strings := [COUNT]string{"abc", "ABC", "123", "一二三"}

	//	定义指针数组
	var ptr [COUNT]*string
	fmt.Printf("%T,%v \n", ptr, ptr)

	// 为指针数组赋值
	for i := 0; i < COUNT; i++ {
		ptr[i] = &strings[i]
	}

	fmt.Printf("%T,%v \n", ptr, ptr)
	// 输出数组中的第一个值
	fmt.Println(ptr[0])
	//	根据数组元素的每个地址获取该地址所指向的元素的数值
	for i := 0; i < COUNT; i++ {
		fmt.Printf("a[%d] = %s \n", i, *ptr[i])
	}
}
