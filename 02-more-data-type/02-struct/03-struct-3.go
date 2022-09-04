package main

import "fmt"

// Vertex3
//结构体字段可以通过结构体指针来访问。
//如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。
//不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
type Vertex3 struct {
	X int
	Y int
}

func main() {

	v := Vertex3{1, 2}
	p := &v          // 取变量v的地址，赋值给p这个指针，完整的写法：*p.X
	fmt.Println(p.X) // 完整的写法：*p.X
	fmt.Println(p.Y) // 完整的写法：*p.Y
	fmt.Println(v.X)
	fmt.Println(v.Y)
	p.X = 1e9
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(v.X)
	fmt.Println(v.Y)
}
