package main

import "fmt"

/**
Go语言源文件默认采用的Unicode字符集。Unicode字符集是目前市面上最流行的字符集，几乎囊括了所有主流非ASCII字符（包括中文字符）。
Go字符串的每个字符都是一个Unicode字符，并且这些Unicode字符是以UTF-8编码格式存储在内存当中的。
*/
func main() {
	// 中文字符  Unicode码点  UTF8编码
	// 中       U+4E2D 	    E4B8AD
	// 国       U+56FD      E59BBD
	// 欢       U+6B22      E6ACA2
	// 迎       U+8FCE      E8BF8E
	// 您       U+60A8      E682A8

	s := "中国欢迎您"
	rs := []rune(s)
	bs := []byte(s)

	for i, r := range rs {
		var utf8byte []byte
		for j := i * 3; j < (i+1)*3; j++ {
			utf8byte = append(utf8byte, bs[j])
		}

		fmt.Printf("%s => %X => %X\n", string(r), r, utf8byte)
	}
}
