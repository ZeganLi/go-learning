package main

import "fmt"

/**
字符串一单定义，则在整个声明周期内不可改变
*/
func main() {

	s := "hello"
	fmt.Println("original string:", s)

	bs := []byte(s)
	bs[0] = 't'
	fmt.Println("slice:", string(bs))

	fmt.Println("after reslice, the original string is	:", string(s))
}
