package main

import "fmt"

//Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
//而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。具体defer.png图所示

//是执行到函数体尾部返回，还是在某个错误处理分支显式调用return返回，抑或出现panic，已经存储到deferred函数栈中的函数都会被调度执行
func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5

}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
