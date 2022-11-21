package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

//接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法
func main() {
	myFloat := MyFloat(-math.Sqrt2)
	fmt.Println(myFloat)
}
