package main

import "fmt"

/**
slice 值传递与引用传递对比
*/
func main() {
	a := []int{1, 2, 3, 4}
	fmt.Printf("1. 变量a的内存地址是：%p,值：%v \n\n", &a, a) // 变量a的内存地址，变量的值
	fmt.Printf("切片型变量a的内存地址是：%p, \n\n", a)         // 数组的内存地址，即a的值的内存地址

	// 传值
	changeSliceVal(a)
	fmt.Printf("2. changeArrayVal函数调用后：变量a的内存地址是：%p,值：%v \n\n", &a, a) // 内容数据已发生变化 [99 2 3 4]

	// 传引用
	changeSlicePtr(&a)
	fmt.Printf("3. changeArrayPtr函数调用后：变量a的内存地址是：%p,值：%v \n\n", &a, a)

	fmt.Println("通过观察1、2、3步骤，可知，变量a的地址没有变化，但数据内容已经通过调用函数，传入引用的方式发生了变化")
}

func changeSlicePtr(a *[]int) {
	fmt.Printf("-----------------------------changeArrayVal函数内：值参数a的内存地址是：%p，值为：%v \n", &a, a)

	(*a)[1] = 250
}

func changeSliceVal(a []int) {
	fmt.Printf("-----------------------------changeArrayVal函数内：值参数a的内存地址是：%p，值为：%v \n", &a, a) // 复制的一份内存地址
	fmt.Printf("-----------------------------changeArrayVal函数内：值参数a的内存地址是：%p \n", a)           // 值未变化
	a[0] = 99
}
