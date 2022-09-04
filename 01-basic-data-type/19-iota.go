package main

import "fmt"

// 枚举常量代表了一类现实的需求
// 1、有限数量标识符构成的集合
// 2、注重类型安全

// 一、一般的枚举常量
//const (
//	Sunday    = 0
//	Monday    = 1
//	Tuesday   = 2
//	Wednesday = 3
//	Thursday  = 4
//	Friday    = 5
//	Saturday  = 6
//)

// 二、隐式重复前一个非空表达式
const (
	Apple, Banana = 11, 22
	Strawberry, Grape
	Pear, Watermelon
)

// 常量定义的后两行没有显式给予初始赋值，Go编译器将为其隐式使用第一行的表达式，这样上述定义等价于
// const (
//  Apple, Banana = 11, 22
//  Strawberry, Grape = 11, 22
//  Pear, Watermelon = 11, 22
//)

// 三、iota。iota是Go语言的一个预定义标识符，它表示的是const声明块（包括单行声明）中每个常量所处位置在块中的偏移值（从零开始）。
// 同时，每一行中的iota自身也是一个无类型常量，可以像无类型常量那样自动参与不同类型的求值过程，而无须对其进行显式类型转换操作。

// $GOROOT/src/sync/mutex.go (go 1.12.7)
const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift      = iota
	starvationThresholdNs = 1e6
)

// 这是一个很典型的诠释iota含义的例子，我们逐行来看。
// mutexLocked = 1 << iota：这里是const声明块的第一行，iota的值是该行在const块中的偏移量，
// 因此iota的值为0，我们得到mutexLocked这个常量的值为1 << 0，即1。

// mutexWoken：这里是const声明块的第二行，由于没有显式的常量初始化表达式，根据const声明块的“隐式重复前一个非空表达式”机制，
// 该行等价于mutexWoken = 1 << iota。由于该行是const块中的第二行，因此偏移量iota的值为1，我们得到mutexWoken这个常量的值为1<< 1，即2。

// mutexStarving：该常量同mutexWoken，该行等价于mutexStarving = 1 << iota，由于在该行的iota的值为2，
// 因此我们得到mutexStarving这个常量的值为 1 << 2，即4。

// mutexWaiterShift = iota：这一行的常量初始化表达式与前三行不同，由于该行为第四行，iota的偏移值为3，因此mutexWaiterShift的值就为3。

// 位于同一行的iota即便出现多次，其值也是一样的：
// const (
// 	Apple, Banana     = iota, iota + 10 // 0, 10 (iota = 0)
//	Strawberry, Grape                   // 1, 11 (iota = 1)
//	Pear, Watermelon                    // 2, 12 (iota = 2)
// )

const (
	PI = 3.1415926
	PI_2
	PI_4 = iota
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {

	fmt.Println(Apple, Banana, Strawberry, Grape, Pear, Watermelon) // 11 22 11 22 11 22

	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexStarving, mutexWaiterShift, starvationThresholdNs)

	fmt.Println(PI, PI_2, PI_4)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	fmt.Printf("%T \r\n", Sunday)
	fmt.Printf("%T \r\n", Wednesday)

	print(Weekday(Monday))
}
