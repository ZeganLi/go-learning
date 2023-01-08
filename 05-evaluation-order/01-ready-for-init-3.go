package main

import "fmt"

// 当多个变量在声明语句左侧且右侧为单一表达式时的表达式求值情况
var (
	n    = m
	j, m = fk()
	k    = 3
)

func fk() (int, int) {
	k++
	return k, k + 1
}

func main() {

	fmt.Println(n, j, m, k)
}
