package main

import (
	"errors"
	"fmt"
)

// 函数作为Go中的"一等公民"，Go中的函数可以像普通整型值那样被创建和使用。
// 有名称的函数
func hashNameFunc() {

}
func main() {

	// 匿名函数因为没有名字，不能像普通函数那样调用，所以需要保存到变量中，或是立即执行
	// 1、匿名函数立即执行
	func() {
		fmt.Println("这是一个无参无返回值的匿名函数")
	}() // 后面加上"()"表示立即调用

	// 2、将匿名函数赋值给一个变量
	var f = func(name string) {
		fmt.Printf("hello : %s", name)
	}
	f("张三")

	addValue := add(1, 2)
	fmt.Println(addValue)

	subValue := sub(10, 8)
	fmt.Println(subValue)

	// 3、函数作为参数
	addOp := calc(10, 20, add)
	fmt.Println(addOp)

	subOp := calc(10, 20, sub)
	fmt.Println(subOp)

	// 4、函数作为返回值
	f2, err := do("+")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f2(100, 200))
	}

}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}
