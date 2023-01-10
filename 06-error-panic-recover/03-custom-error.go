package main

import (
	"fmt"
	"time"
)

// 自定义错误的实现步骤如下。
// • 定义一个结构体，表示自定义错误的类型。
// • 让自定义错误类型实现error接口：Error() string。
// • 定义一个返回error的函数。根据程序实际功能而定

// MyError 定义一个结构体，表示自定义错误的类型
type MyError struct {
	When time.Time
	What string
}

// 实现error接口的方法
func (m MyError) Error() string {
	return fmt.Sprintf("%v : %v", m.When, m.What)
}

func getArea(width, length float64) (float64, error) {
	errorInfo := ""
	if width < 0 && length < 0 {
		errorInfo = fmt.Sprintf("长度: %v,宽度: %v,均为负数", length, width)
	} else if length < 0 {
		errorInfo = fmt.Sprintf("长度: %v 出现负数", length)
	} else if width < 0 {
		errorInfo = fmt.Sprintf("宽度: %v 出现负数", width)
	}

	if errorInfo != "" {
		return 0, MyError{time.Now(), errorInfo}

	} else {
		return width * length, nil
	}
}

func main() {
	res, err := getArea(-4, -5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
