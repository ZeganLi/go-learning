package main

import "fmt"

/**
数组是值类型

当它们被分配给一个新变量时，会将原始数组复制出一份分配给新变量。因此对新变量进行更改，原始数组不会有反应。
将数组作为函数参数进行传递，它们将通过值传递，原始数组依然保持不变。
*/
func main() {

	coutrys1 := [...]string{"USA", "China", "India", "Germany", "France"}
	coutrys2 := coutrys1

	coutrys2[0] = "Japanese"

	fmt.Println("coutrys1 :", coutrys1) // coutrys1 : [USA China India Germany France]
	fmt.Println("coutrys2 :", coutrys2) // coutrys2 : [Japanese China India Germany France]

	coutrys3 := &coutrys2[0]
	*coutrys3 = "泰国"
	fmt.Printf("coutrys3的值为 = %v\n", *coutrys3)
	fmt.Printf("coutrys1的地址值为 = %v \n", coutrys1)
	fmt.Printf("coutrys2的地址值为 = %v \n", coutrys2)
	fmt.Printf("coutrys3的地址值为 = %v \n", coutrys3)
}
