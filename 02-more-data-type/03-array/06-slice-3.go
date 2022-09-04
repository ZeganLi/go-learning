package main

import (
	"fmt"
)

/**
切片文法
切片文法类似于没有长度的数组文法。

这是一个数组文法：

[3]bool{true, true, false} ** 数组需要指定长度
下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

[]bool{true, true, false}  ** 切片没有
*/

type s2 struct {
	i int
	b bool
}

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// 官网示例
	//s := []struct {
	//	i int
	//	b bool
	//}{
	//	{2, true},
	//	{3, false},
	//	{5, true},
	//	{7, true},
	//	{11, false},
	//	{13, true},
	//}

	// 改写示例1
	//s2s := [6]s2{}
	//
	//for i := range s2s {
	//	s2s[i] = s2{q[i], r[i]}
	//}
	//fmt.Println(s)
	//fmt.Println(s2s)

	// 改写示例2
	s3s := make([]s2, 0)
	for i := range q {
		s3s = append(s3s, s2{q[i], r[i]})
	}

	fmt.Println(s3s)

}
