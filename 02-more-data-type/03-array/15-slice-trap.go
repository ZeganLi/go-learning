package main

import (
	"io/ioutil"
	"regexp"
)

// 可能的陷阱
var digitRegexp = regexp.MustCompile("[0-9]+")

// FindDigits 函数加载整个文件到内存，然后搜索第一个连续的数字，最后结果以切片方式返回。
// 这段代码的行为和描述类似，返回的 []byte 指向保存整个文件的数组。
// 因为切片引用了原始的数组， 导致 GC 不能释放数组的空间；
// 只用到少数几个字节却导致整个文件的内容都一直保存在内存里。
func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

// CopyDigits 要修复整个问题，可以将感兴趣的数据复制到一个新的切片中：
func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
