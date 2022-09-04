package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型转string，有两种方法
// 1. fmt.Sprint([int,bool,float...])
// 2. strconv.[FormatInt,FormatBool,FormatFloat...]
func main() {

	basicInt := 123456.789
	sprint := fmt.Sprint(basicInt)
	formatInt := strconv.FormatInt(int64(basicInt), 10)
	fmt.Printf("basicInt 的数据类型是%T,", basicInt)
	fmt.Printf("basicInt 的数据类型是%T,", formatInt)
	fmt.Printf("basicInt 转换过后的数据类型是%T\n", sprint)
	fmt.Printf("basicInt 转换过后的数据类型是%T\n", formatInt)

	basicFloat := 3.14159
	basicFloatStr := fmt.Sprint(basicFloat)
	fmt.Printf("basicFloat 的数据类型是%T,", basicFloat)
	fmt.Printf("basicFloat 转换过后的数据类型是%T", basicFloatStr)

}
