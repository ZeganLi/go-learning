package main

import "fmt"

/**
多重指针
*/
func main() {

	var v int8 = 127

	var p1 *int8
	p1 = &v

	var p2 **int8
	p2 = &p1

	var p3 ***int8
	p3 = &p2

	var p4 ****int8
	p4 = &p3

	fmt.Printf("一重指针p1的类型：%T,p1的值：%d\n", p1, p1)
	fmt.Printf("二重指针p2的类型：%T,p2的值：%d\n", p2, p2)
	fmt.Printf("三重指针p3的类型：%T,p3的值：%d\n", p3, p3)
	fmt.Printf("四重指针p4的类型：%T,p4的值：%d\n", p4, p4)

	fmt.Printf("一重指针p1的类型：%T,p1的值：%d\n", p1, *p1)
	fmt.Printf("二重指针p2的类型：%T,p2的值：%d\n", p2, *p2)
	fmt.Printf("三重指针p3的类型：%T,p3的值：%d\n", p3, *p3)
	fmt.Printf("四重指针p4的类型：%T,p4的值：%d\n", p4, *p4)
}
