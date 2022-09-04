package main

import "fmt"

// 如果包级别的变量中使用了空变量_，空变量也会得到Go编译器一视同仁的对待。如下：
var (
	x = y + z
	y = fun()
	_ = fun()
	z = fun()
	h = 3
)

func fun() int {
	h++
	return h
}

func main() {
	fmt.Println(x, y, z, h)
}
