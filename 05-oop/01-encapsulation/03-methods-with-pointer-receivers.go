package main

import (
	"fmt"
	"math"
)

type VertexPointer struct {
	X, Y float64
}

func (v *VertexPointer) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *VertexPointer, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f

}

func (v VertexPointer) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v *VertexPointer) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

func main() {
	//vertexPointer := VertexPointer{3, 4}
	//vertexPointer.Scale(10)
	//fmt.Println(vertexPointer)
	//
	//vertexPointer.Scale(10)
	//fmt.Println(vertexPointer)

	vertexPointer := VertexPointer{3, 4}

	//Scale(vertexPointer, 10)  // 报错，调用指针函数时，必须传递一个指针
	Scale(&vertexPointer, 10) // 编译正确

	//Abs(vertexPointer)
	Abs(&vertexPointer)

	vertexPointer.Abs() // 当调用一个值的方法时，即可以使用值调用，也可以使用指针调用
	p := &vertexPointer
	p.Abs()

	fmt.Println(p)

	//所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用
}
