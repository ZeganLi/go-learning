package main

import (
	"fmt"
	"math"
)

/*
在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。例如，Pizza 就是个已导出名，Pi 也同样，它导出自 math 包。
pizza 和 pi 并未以大写字母开头，所以它们是未导出的。
在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。

-- 给外界提供的所有的函数、常量、字段等，都应该以大写字母开头。
*/
func main() {
	fmt.Println(math.Pi)
	//fmt.Println(math.pi) // .\05-exported-names.go:14:14: cannot refer to unexported name math.pi
}
