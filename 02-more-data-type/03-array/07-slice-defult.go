package main

import "fmt"

/**
切片的默认行为
在进行切片时，你可以利用它的默认行为来忽略上下界。

切片下界的默认值为 0，上界则是该切片的长度。

对于数组

var a [10]int
来说，以下切片是等价的：

a[0:10]
a[:10]
a[0:]
a[:]
*/
func main() {

	//可以通过设置下限及上限来设置截取切片：[lower-bound:upper-bound]

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[:]
	fmt.Println(s) // [2 3 5 7 11 13]

	// ！！！切片是在原有的基础上进行的再切片，会改变原有的切片的值！！！
	s = s[1:4]
	fmt.Println(s) // [3 5 7]

	s = s[:2]
	fmt.Println(s) // [3 5]

	s = s[1:]
	fmt.Println(s) // [5]
}
