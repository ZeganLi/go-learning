package main

import "fmt"

//切片的自动扩容：
// 1、当底层数组的容量不满足新加的元素的时候，会触发自动扩容。
// 2、扩容时，len + n(新添加的元素个数) cap = old cap * 2
func main() {
	var s []int // s被赋予nui值
	s = append(s, 11)
	fmt.Printf("s的长度为: %d,s的容量为 %d \n", len(s), cap(s))

	s = append(s, 22)
	fmt.Printf("s的长度为: %d,s的容量为 %d \n", len(s), cap(s))

	s = append(s, 33)
	fmt.Printf("s的长度为: %d,s的容量为 %d \n", len(s), cap(s))

	s = append(s, 44)
	fmt.Printf("s的长度为: %d,s的容量为 %d \n", len(s), cap(s))

	s = append(s, 55)
	fmt.Printf("s的长度为: %d,s的容量为 %d \n", len(s), cap(s))

	s = append(s, 66)
	fmt.Printf("s的长度为: %d,s的容量为 %d \n", len(s), cap(s))

	s = append(s, 77)
	fmt.Printf("s的长度为: %d,s的容量为 %d \n", len(s), cap(s))

	s = append(s, 88)
	fmt.Printf("s的长度为: %d,s的容量为 %d \n", len(s), cap(s))

	s = append(s, 99)
	fmt.Printf("s的长度为: %d,s的容量为 %d \n", len(s), cap(s))

	fmt.Println(s)
}
