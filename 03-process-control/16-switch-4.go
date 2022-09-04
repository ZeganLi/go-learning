package main

import "fmt"

/**
switch 后可以匹配多个表达式
*/
func main() {
	var score int
	fmt.Println("请输入成绩：")
	fmt.Scanf("%d", &score)
	switch {
	case score >= 90 && score <= 100: // switch的值与case的类型必须一致
		fmt.Println("优秀")

	case score < 90 && score >= 70: // case后可以跟多个值，匹配其中任意一个
		fmt.Println("优良")
		fallthrough // 穿透，会执行到下一个case

	case score >= 60 && score < 70:
		fmt.Println("及格")

	default:
		fmt.Println("不及格")
	}
}
