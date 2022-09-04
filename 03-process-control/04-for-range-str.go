package main

import "fmt"

/**
当使用 for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
注意是副本，不是原来的
i和v在整个for range过程中被重用
*/
func main() {

	var str = "hello,北京"
	for i, v := range str {
		//fmt.Printf("i的值是%v,%c\n", i, v)
		fmt.Println("i的值是：", i, "v的地址值：", &v) // i的值是从0——9共十个字节 v的地址值一直不变,也就相当于是重用了i
		//fmt.Println(i, v) // 当输出的是字符串的时候，默认是按照ASCII码值进行的输出，如需要输出ASCII码对应的值，则使用%c
	}

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // hello,åäº¬ 乱码 使用fori的形式对字符串进行遍历时，默认是按照字节进行输出
	}

	for _, v := range str {
		fmt.Printf("%c", v) // hello,北京 使用for range的形式对字符串进行遍历时，默认是按照字符进行输出
	}
}
