package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func main() {
	wc.Test(WordCount)
	wc.Test(FieldsTest)
}

func FieldsTest(s string) map[string]int {
	// 根据 unicode.IsSpace 的定义，Fields 围绕一个或多个连续空白字符的每个实例拆分字符串 s，
	// 如果 s 仅包含空白，则返回 s 的子字符串切片或空切片。
	split := strings.Fields(s)

	var m = make(map[string]int)

	for _, v := range split[:] {
		_, ok := m[v]

		if ok {
			m[v]++

		} else {
			m[v] = 1
		}
	}
	return m
}

// WordCount
// 实现 WordCount。
// 它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。
// 函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。
func WordCount(s string) map[string]int {
	split := strings.Split(s, " ")

	var m = make(map[string]int)

	for _, v := range split[:] {
		_, ok := m[v]

		if ok {
			m[v]++

		} else {
			m[v] = 1
		}
	}
	return m
}
