package main

import "fmt"

// panic，让当前的程序进入恐慌，中断程序的执行。panic()是一个内建函数，可以中断原有的控制流程
// 情况下，向程序使用方报告错误状态的方式可以是返回一个额外的error类型值。但是，当遇到不可恢复的错误状态时，如数组访问越界、空指针引用等，
// 这些运行时错误会引起panic异常。这时，上述错误处理方式显然就不适合了。
// 需要注意的是，不应通过调用panic()函数来报告普通的错误，而应该只把它作为报告致命错误的一种方式。当某些不应该发生的场景发生时调用panic()。
func main() {
	TestA1()
	TestB1(101)
	TestC1()
}

func TestA1() {
	fmt.Println("func TestA()")
}

func TestB1(x int) {
	defer func() {
		msg := recover()
		fmt.Println("发生异常: ", msg)
	}()

	var a [100]int
	a[x] = 1000 // x值超出101时，数组越界
}

func TestC1() {
	fmt.Println("func TestC()")

}
