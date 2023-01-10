package main

import "fmt"

// recover()可以让进入恐慌流程的Goroutine恢复过来并重新获得流程控制权。
// 需要注意的是，recover()让程序恢复，必须在延迟函数中执行。换言之，recover()仅在延迟函数中有效。
// 在正常的程序运行过程中，调用 recover()会返回 nil，并且没有其他任何效果。如果当前的Goroutine陷入恐慌，调用recover()可以捕获panic()的输入值，使程序恢复正常运行。
// 开发者应该把它作为最后的手段来使用，换言之，开发者的代码中尽量少有或者没有panic异常
func main() {
	TestA2()
	TestB2()
	TestC2()
	fmt.Println("main over")

}

func TestA2() {
	fmt.Println("这是 funcA2")
}

func TestB2() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("恢复啦, 获取到recover()的值: %v", msg)
		}
	}()

	fmt.Println("这是 funcB2")

	for i := 0; i < 10; i++ {
		fmt.Println("i : ", i)
		if i == 5 {
			panic("funcB2恐慌了")
		}
	}
}

func TestC2() {
	defer func() {
		fmt.Println("执行TestC2的延迟函数")
		msg := recover()
		fmt.Println("恢复啦, 获取到recover()的值: %v", msg)
	}()

	fmt.Println("这是TestC2")
	panic("TestC2 发起panic")
}
