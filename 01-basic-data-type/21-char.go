package main

import "fmt"

func main() {

	var c = 'a'

	fmt.Printf("c 的类型是=%T,c的值是%c \n", c, c) // %c 相应Unicode码点所表示的字符
	fmt.Println(c)                          // 直接输出的是ASCII码

}
