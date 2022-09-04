package main

import (
	"fmt"
)

/*
--数据类型转换，必须显示的声明需要转换的值，和类型

表达式 T(v) 将值 v 转换为类型 T。
一些关于数值的转换：

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
或者，更加简单的形式：

i := 42
f := float64(i)
u := uint(f)
*/

func main() {
	var x, y int = 3, 4
	var f float64 = float64(x*x + y*y)
	//var f float64 = math.Sqrt(x*x + y*y) //cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt
	fmt.Println(f)
	var z float32 = float32(f)
	//var z uint = f // cannot use f (type float64) as type uint in assignment

	//g := 0.867 + 0.5i // complex128

	fmt.Println(x, y, z)

	fmt.Println(2i)
}
