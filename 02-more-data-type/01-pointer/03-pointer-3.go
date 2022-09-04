package main

import "fmt"

// 取地址符是 &
// 取指针所对应的值是 *ptr
func main() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a) // c0000160a8

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip) // c0000160a8

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip) // 20

	// 修改ip指针指向的值
	*ip = 40
	fmt.Printf("*ip新的变量的值为：%d\n", *ip) // 40
	fmt.Printf("ip新的地址的值为：%d\n", ip)   // 824633811112
	fmt.Printf("&ip新的地址的值为：%d\n", &ip) // 824633745448
}
