package main

import (
	"fmt"
)

// Vertex2 一个结构体（struct）就是一组字段（field）
type Vertex2 struct {
	X int
	Y int
}

// 结构体字段使用点号来访问
func main() {
	vertex := Vertex2{12, 24}

	fmt.Println(vertex.X)
	vertex.X = vertex.X << 1
	fmt.Println(vertex.X)

	fmt.Println(vertex.Y)
	vertex.Y = vertex.Y << 1
	fmt.Println(vertex.Y)
}
