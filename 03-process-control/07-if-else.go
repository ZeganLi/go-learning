package main

import (
	"fmt"
	"math"
)

func pow1(x, n, lim float64) float64 {
	// 在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}
func main() {
	// fmt.Println 调用开始前，两次对 pow 的调用均已执行并返回其各自的结果。
	fmt.Println(pow1(3, 2, 10))
	fmt.Println(pow1(3, 3, 20))
}
