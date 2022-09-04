package main

import "fmt"

/**
nil 切片
切片的零值是 nil。

nil 切片的长度和容量为 0 且没有底层数组。 ** 相当于 切片 -->> nil
*/
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s)) // len()和cap的返回值都是int类型，int类型默认值为0，所以切片的返回值也是0
	if s == nil {
		fmt.Println("nil!")
	}
}
