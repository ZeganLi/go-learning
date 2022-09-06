package main

import (
	"fmt"
	_ "go-tour/04-function-closure/pkg1" // 即使忽略了，没有使用，也会触发其他包内变量的初始化等操作
)

var (
	_  = constInitCheck() // 不会输出，相当于没有引用
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

const (
	c1 = "c1"
	c2 = "c2"
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("main: const c1 init")
	}
	if c1 != "" {
		fmt.Println("main: const c2 init")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("main: var %s init\n", name)
	return name
}
func main() {
	fmt.Println("main: ")

	//Go运行时按照pkg2→pkg1→pkg3→main的包顺序以及在
	//包内常量→变量→init函数的顺序进行初始化。
}

func init() {
	fmt.Println("main: init")
}
