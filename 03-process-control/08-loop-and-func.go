package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	var temp float64
	for i := 0; i < 9; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("第%d次循环,求得x的平方根为：%d\n", i+1, z)

		fmt.Printf("math:%d\n", math.Sqrt(x))

		if temp == z {
			return z
		} else {
			temp = z
		}
	}

	return 0
}
func main() {
	fmt.Println(Sqrt(2))
}
