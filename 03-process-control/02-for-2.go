package main

import "fmt"

//初始化语句和后置语句是可选的。
func main() {
	sum := 0
	// Go 中的"while"
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
