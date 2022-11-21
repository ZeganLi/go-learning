package main

import (
	"fmt"
	"math"
)

/**
go中的方法的本质是：方法就是一类带特殊的 <接收者> 参数的函数。
*/

type Vertex struct {
	X, Y float64
}

// Abs 在func关键字和Abs方法名之间,写上指定的类型
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs1 这种方式与Abs一样，一个是使用struct来调用，一个是通过函数传入具体的类型的方式来调用
func Abs1(v Vertex) float64 {
	return math.Abs(v.X*v.X + v.Y*v.Y)
}

// Scale 可以来传递一个指针类型的struct，此种方式可以修改指针对象的值，不是指针的方式，是传入变量类型的一个副本，无法修改原对象的值
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	//abs1 := Abs1(v)
	//fmt.Println(v.Abs())
	//fmt.Println(abs1)

	v.Scale(10)
	fmt.Println(v.Abs())
}
