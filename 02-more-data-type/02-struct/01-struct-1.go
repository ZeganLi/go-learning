package main

import "fmt"

// Vertex 一个结构体（struct）就是一组字段（field）
type Vertex struct {
	X int
	Y int
}

// VertexPtr 函数参数也可使用自定义类型的指针类型
func VertexPtr(vertex *Vertex) {
	vertex.Y++
	vertex.X--
}

func main() {
	// 基本初始化 初始构造器
	vertex := Vertex{10, 20}

	fmt.Println(vertex)
	var verPtr *Vertex = &vertex
	VertexPtr(verPtr)
	fmt.Println(verPtr.X, verPtr.Y)

	// 高级初始化 指定字段和值 Go推荐使用field:value的复合字面值形式对struct类型变量进行值构造，
	// 这种值构造方式可以降低结构体类型使用者与结构体类型设计者之间的耦合，这也是Go语言的惯用法。
	vertex2 := Vertex{X: 20, Y: 30}
	fmt.Println(vertex2)
	var verPtr2 *Vertex = &vertex2
	VertexPtr(verPtr2)
	fmt.Println(verPtr2.X, verPtr2.Y)

	//复合字面值作为结构体值构造器的大量使用，使得即便采用类型零值时我们也会使用字面值构造器形式
	vertex3 := Vertex{} // 常用

	//而较少使用new这一个Go预定义的函数来创建结构体变量实例：
	vertex4 := new(Vertex) // 较少使用

	fmt.Println(vertex3)
	fmt.Println(vertex4)
	//不允许将从其他包导入的结构体中的未导出字段作为复合字面值中的field
}
