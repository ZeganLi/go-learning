package main

import (
	"fmt"
	"time"
)

func main() {
	//defraud1()
	//defraud2()
	//defraud3()
	//defraud4()
	defraud5()
}

func defraud5() {
	//原切片a在for range的循环过程中被附加了两个元素6和7，其len由5增加到7，但这对于r却没有产生影响。
	//其原因就在于a的副本a'的内部表示中的len字段并没有改变，依旧是5，因此for range只会循环5次，
	//也就只获取到a所对应的底层数组的前5个元素。
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	fmt.Println("a = ", a)
	// 长度的变化
	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}

		r = append(r, v)
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
func defraud4() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("defraud1 result : ")
	fmt.Println(a) // 1 2 3 4 5

	// 使用基于a的切片
	for i, v := range a[:] {
		if i == 0 {
			// 对a的修改，会导致接下来的循环的v值的变化
			a[1] = 12
			a[2] = 13
		}

		//fmt.Println("循环内值的地址值是：", &v)
		//fmt.Println("循环外a数组的地址值为：", &a[i])
		r[i] = v
	}

	fmt.Println(r) // 1 12 13 4 5
	fmt.Println(a) // 1 12 13 4 5
}
func defraud3() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("defraud1 result : ")
	fmt.Println(a) // 1 2 3 4 5

	// 与defraud2不同的是使用a数组的指针
	for i, v := range &a {
		if i == 0 {
			// 对a的修改，并不会导致接下来的循环的v值的变化
			a[1] = 12
			a[2] = 13
		}

		//fmt.Println("循环内值的地址值是：", &v)
		//fmt.Println("循环外a数组的地址值为：", &a[i])
		r[i] = v
	}

	fmt.Println(r) // 1 12 13 4 5
	fmt.Println(a) // 1 12 13 4 5
}
func defraud2() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("defraud1 result : ")
	fmt.Println(a) // 1 2 3 4 5

	// 参与range循环的a 和 上面的a不是同一块儿内存区域。参与循环的是真正a的复制。因此如果是对一个大数组进行range，会导致数组的复制
	for i, v := range a {
		if i == 0 {
			// 对a的修改，并不会导致接下来的循环的v值的变化
			a[1] = 12
			a[2] = 13
		}

		fmt.Println("循环内值的地址值是：", &v)
		fmt.Println("循环外a数组的地址值为：", &a[i])
		r[i] = v
	}

	fmt.Println(r) // 1 2 3 4 5
	fmt.Println(a) // 1 12 13 4 5
}

// 1、for range i和v在整个range过程中被重用
//goroutine中输出的i、v值都是for range循环结束后的i、v最终值，
//而不是各个goroutine启动时的i、v值。这是因为goroutine执行的闭包函数引用了它的外层包裹函数中的变量i、v，
//这样变量i、v在主goroutine和新启动的goroutine之间实现了共享。而i、v值在整个循环过程中是重用的，仅有一份。
//在for range循环结束后，i = 4，v = 5，因此各个goroutine在等待3秒后进行输出的时候，输出的是i、v的最终值。
func defraud1() {
	var m = [...]int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v) // 4 5
		}()
	}

	time.Sleep(time.Second * 10)
}
