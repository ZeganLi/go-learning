package main

import "fmt"

func main() {
	//test1()
	//test2()
	//test3()

}

func test3() {
	// 由于Go string是不可变的，因此如果两个字符串的长度不相同，那么无须比较具体字符串数据即可断定两个字符串是不同的。
	// 如果长度相同，则要进一步判断数据指针是否指向同一块底层存储数据。
	// 如果相同，则两个字符串是等价的；如果不同，则还需进一步比对实际的数据内容。
	// ==
	s1 := "世界和平"
	s2 := "世界" + "和平"
	fmt.Println(s1 == s2) // true

	// !=
	s1 = "Go"
	s2 = "C"
	fmt.Println(s1 != s2) // true

	// < 和 <=
	s1 = "12345"
	s2 = "23456"
	fmt.Println(s1 < s2)  // true
	fmt.Println(s1 <= s2) // true

	// > 和 >=
	s1 = "12345"
	s2 = "123"
	fmt.Println(s1 > s2)  // true
	fmt.Println(s1 >= s2) // true
}

func test2() {
	//支持通过+/+=操作符进行字符串连接
	s := "hello "
	s2 := s + "world "
	s2 += "go"
	fmt.Println(s2)
}

func test1() string {
	// Go string类型支持“零值可用”的理念。Go字符串无须像C语言中那样考虑结尾'\0'字符，因此其零值为""，长度为0
	var s string
	fmt.Println(s) // s = ""
	//获取长度的时间复杂度是O(1)级别
	fmt.Println(len(s)) // 0
	return s
}
