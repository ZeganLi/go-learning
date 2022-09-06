package main

import (
	"fmt"
	"strings"
)

func main() {

	//adderTest1()
	//adderTest2()
	//suffixTest()
	add, sub := calc2(100)
	println(add(50)) // 150
	println(sub(50)) // 100
	// 牢记——闭包=函数+引用环境
}

func suffixTest() {
	suffixFunc := makeSuffixFunc(".avi")
	s := suffixFunc("激情燃烧的岁月.mp4")
	fmt.Println(s)

	f := makeSuffixFunc(".jpg")
	s2 := f("苍井空.jpg")
	fmt.Println(s2)
}

func adderTest2() {
	f := adder2(10)
	fmt.Println(f(10)) // 20
	fmt.Println(f(20)) // 40
	fmt.Println(f(30)) // 70

	f2 := adder2(100)   // x 重新计数
	fmt.Println(f2(40)) // 140
	fmt.Println(f2(50)) // 190
}

func adderTest1() {
	//闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，"闭包=函数+引用环境"。
	f := adder()
	fmt.Println(f(10)) // 10
	fmt.Println(f(20)) // 30
	fmt.Println(f(30)) // 60

	f2 := adder()       // x 重新计数
	fmt.Println(f2(40)) // 40
	fmt.Println(f2(50)) // 90

	//变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效
}

func adder() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(y int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(s string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 返回多个闭包
func calc2(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub

}
