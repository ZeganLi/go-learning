package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 映射的零值为 nil 。nil 映射既没有键，也不能添加键
var m map[string]Vertex // 默认值为nil，不符合golang的"零值可用"
func main() {
	//make 函数会返回给定类型的映射，并将其初始化备用
	m = make(map[string]Vertex)

	m["ball labs"] = Vertex{
		Lat:  40.68433,
		Long: -74.39967,
	}

	m["beijing"] = Vertex{
		Lat:  33.7894,
		Long: 78.22121,
	}

	// 若顶级类型只是一个类型名，你可以在文法的元素中省略它。
	var m = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}
