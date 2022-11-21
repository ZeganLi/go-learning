package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 原始string
	var s string = "hello"
	fmt.Println("original string:", s) // 试图通过unsafe指针改变原始string
	modifyString(&s)
	fmt.Println(s)
}

/**
对string的底层的数据存储区仅能进行只读操作，一旦试图修改那块区域的数据，便会得到SIGBUS的运行时错误
*/
func modifyString(s *string) {
	// 取出第一个8字节的值
	p := (*uintptr)(unsafe.Pointer(s))

	// 获取底层数组的地址
	var array *[5]byte = (*[5]byte)(unsafe.Pointer(*p))
	var len *int = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Sizeof((*uintptr)(nil))))

	for i := 0; i < (*len); i++ {
		fmt.Printf("%p => %c\n", &((*array)[i]), (*array)[i])
		p1 := &((*array)[i])
		v := *p1
		*p1 = v + 1 //try to change the character
	}
}
