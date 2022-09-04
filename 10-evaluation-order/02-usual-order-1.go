package main

import "fmt"

// 普通求值顺序

// 按照从左到右的顺序，先对等号左侧表达式操作数中的函数进行调用求值，因此第一个是y[f()]中的f()。
// 接下来是等号右侧的表达式。第一个函数是g()，但g()依赖其参数的求值，其参数列表依然可以看成是一个多值赋值操作，
// 其涉及的函数调用顺序从左到右依次为h()、i()、j()、<-c，这样该表达式操作数函数的求值顺序即为h() -> i() -> j() -> c取值操作 -> g()。
// 最后还剩下末尾的k()，因此该语句中函数以及channel操作的完整求值顺序是：f() -> h() -> i() -> j() -> c取值操作 -> g() ->k()。

func main() {
	var y = []int{11, 12, 13}
	var x = []int{21, 22, 23}

	var c chan int = make(chan int)
	go func() { c <- 1 }()
	y[f1()], _ = g1(h1(), i1()+x[j1()], <-c), k1()
}

func f1() int {
	fmt.Println("calling f")
	return 1
}

func g1(a, b, c int) int {
	fmt.Println("calling g")
	return 2
}
func h1() int {
	fmt.Println("calling h")
	return 3
}
func i1() int {
	fmt.Println("calling i")
	return 1
}
func j1() int {
	fmt.Println("calling j")
	return 1
}
func k1() bool {
	fmt.Println("calling k")
	return true
}
