package main

import (
	"fmt"
	"time"
)

//	switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
//	（例如，
//
//	switch i {
//	case 0:
//	case f():
//	}
//	在 i==0 时 f 不会被调用。）
func main() {
	fmt.Println("when's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("ToDay.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two Days.")
	default:
		fmt.Println("Too far away")
	}

}
