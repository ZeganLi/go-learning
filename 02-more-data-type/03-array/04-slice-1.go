package main

import "fmt"

/**
切片
每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。

**
类型 []T 表示一个元素类型为 T 的切片。
** 注意和数组的区别：
类型 [n]T 表示拥有 n 个 T 类型的值的数组。
**

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

a[low : high]
它会选择一个半开区间，包括第一个元素，但排除最后一个元素。

以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：

a[1:4]
*/
func main() {
	// 1、声明一个空切片，该切片默认为nil，长度为0
	var identifier []int
	fmt.Println(identifier == nil) // true
	fmt.Println(len(identifier))   // 0

	// 2、从数组中获取一个切片
	arr := [6]int{2, 3, 5, 7, 11, 13}  // 声明一个数组
	slice := []int{2, 3, 5, 7, 11, 13} // 声明并初始化一个切片

	// 前闭后开，包括1，不包括索引4，到索引3
	var s []int = arr[1:4]   // 对数组进行切片，得到切片之后的数组
	var b []int = slice[1:4] // 对切片再进行切片，得到切片之后的数组
	fmt.Println(s)           // [3 5 7]
	fmt.Println(b)           // [3 5 7]

	// 3、使用make()函数来创建切片
	slices := make([]int, 10, 10) // make([]T, length, capacity)。类型，长度，容量
	fmt.Println(slices)
}
