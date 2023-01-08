package main

import "fmt"

// 当普通求值顺序与包级变量的初始化依赖顺序一并使用时，后者优先级更高，但每个单独表达式中的操作数求值依旧按照普通求值顺

// 根据包变量初始化依赖规则以及普通求值顺序规则对这个例子进行简要分析，把单行的声明语句等价转换为下面的代码，这样看起来更直观。
//（注意：与前面的多个变量在声明语句左侧且右侧为单一表达式时的表达式求值情况不同，这里右侧并非单一表达式。）
var a2, b2, c2 = f2() + v2(), g2(), sqr(u2()) + v2()

// var (
//	 a2 = f2() + v2()
//	 b2 = g2()
//	 c2 = sqr(u2()) + v2()
// )

func main() {
	fmt.Println(a2, b2, c2)
}

func f2() int {
	fmt.Println("calling f")
	return c
}
func g2() int {
	fmt.Println("calling g")
	return 1
}
func sqr(x int) int {
	fmt.Println("calling sqr")
	return x * x
}
func v2() int {
	fmt.Println("calling v")
	return 1
}
func u2() int {
	fmt.Println("calling u")
	return 2
}
