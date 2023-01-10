package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

// 使用error接口
func main() {

	// 异常情况1
	res := math.Sqrt(-100)
	fmt.Println(res)
	res, err := Sqrt(-100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// 异常情况2
	// 变量res并没有被重复创建，自上而下使用的都是一个变量res。因而不能使用:=，直接=赋值即可
	// 但需保证在同一作用域中，如果在不同的作用域，则函数返回的则是一个新创建的变量
	res, err = Divide(100, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}

	// 异常情况3
	f, err := os.Open("/abc.txt")
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(f.Name(), "文件被打开")
	}

}

func Divide(dividee float64, divider float64) (float64, error) {
	if divider == 0 {
		return 0, errors.New("出错: 除数不能为0")
	} else {
		return dividee / divider, nil
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("负数不可以获取平方根")
	} else {
		return math.Sqrt(f), nil
	}
}
