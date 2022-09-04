package main

import "fmt"

/*
常量的声明与变量类似，只不过是使用 const 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 := 语法声明。

Go中常量同Java一样，也是只能赋值一次，不可修改
*/

// 无类型常量
//const Pi = 3.14
//const Pi = "3"
//const Pi = false
const Pi = 0.2

// 有类型常量
const myInt int = 5
const myFloat float64 = 3.1415926
const myString string = "Hello, Gopher"

const (
	a, s = 5, "Hello, Gopher"
)

// 无类型常量使得Go在处理表达式混合数据类型运算时具有较大的灵活性，代码编写也有所简化，我们无须再在求值表达式中做任何显式类型转换了。
// 除此之外，无类型常量也拥有自己的默认类型：
// 无类型的布尔型常量、整数常量、字符常量、浮点数常量、复数常量、字符串常量对应的默认类型分别为
// bool、int、int32(rune)、float64、complex128和string。
func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(myInt)    // 输出：5
	fmt.Println(myFloat)  // 输出：3.1415926
	fmt.Println(myString) // 输出：Hello,Gopher

	var e float64 = Pi + myFloat
	fmt.Println(e)

	// 当常量被赋值给无类型变量、接口变量时，常量的默认类型对于确定无类型变量的类型及接口对应的动态类型是至关重要的。
	// 下面i的类型，根据接口i的类型的变化而动态变化
	n := a
	var i interface{} = a
	fmt.Printf("%T\n", n) // 输出：int
	fmt.Printf("%T\n", i) // 输出：int
	i = s
	fmt.Printf("%T\n", i) // 输出：string

}
