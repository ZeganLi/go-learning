package main

import (
	"bufio"
	"errors"
	"fmt"
)

func main() {

	// 创建error对象方式1
	err1 := errors.New("自己的error错误")
	fmt.Println(err1.Error())
	fmt.Println(err1)
	fmt.Printf("err1 的类型是: %T\n", err1) // *errors.errorString
	fmt.Println("-------------------------------------")

	// 创建error对象方式2
	err2 := fmt.Errorf("错误的类型%d", 10)
	fmt.Println(err2.Error())
	fmt.Println(err2)
	fmt.Printf("err2 的类型是: %T\n", err2) // *errors.errorString
	fmt.Println("-------------------------------------")

	// 创建error对象方式3
	// 当我们在格式化字符串中使用%w时，fmt.Errorf返回的错误值的底层类型为fmt.wrapError
	// 与errorString相比，wrapError多实现了Unwrap方法，这使得被wrapError类型包装的错误值在包装错误链
	// 中被检视（inspect）到：
	var ErrFoo = errors.New("the underlying error")
	err3 := fmt.Errorf("wrap err: %w", ErrFoo)
	errors.Is(err3, bufio.ErrBufferFull) // true (仅适用于Go 1.13及后续版本)
	fmt.Printf("err3 的类型是: %T\n", err3)  //  *fmt.wrapError
	fmt.Println("-------------------------------------")

	// error对象在函数中的使用
	res, err4 := checkAge(-12)
	if err4 != nil {
		fmt.Println(err4.Error())
		fmt.Println(err4)
	} else {
		fmt.Println(res)
	}
}

func checkAge(age int) (string, error) {
	if age < 0 {
		err := fmt.Errorf("您的年龄输入的是: %d,年龄不能为负数", age)
		return "", err
	} else {
		return fmt.Sprintf("您的年龄输入的是: %d ", age), nil
	}

}
