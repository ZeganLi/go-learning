package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// 同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。该语句声明的变量作用域仅在 if 之内。
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}
