package main

import (
	"fmt"
	"runtime"
)

// switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。
// Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case 。
// 实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止。
func main() {

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("windows")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("%s.\n", os)
	}
}
