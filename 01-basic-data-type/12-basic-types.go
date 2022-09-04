package main

import (
	"fmt"
	"math/cmplx"
)

/*
Go 的基本类型有

bool

string	-- 在go中，字符串类型是基本数据类型。并非Java中的引用数据类型

int  int8  int16  int32  int64	-- 有符号数值
uint uint8 uint16 uint32 uint64 uintptr		-- 无符号数据类型，前面带u的为无符号数值类型

byte // uint8 的别名		-- byte 无符号数值，范围 0-255：等于一个字节

rune // int32 的别名		-- rune 有符号的4个字节的int类型数据
    // 表示一个 Unicode 码点

float32 float64		-- 不同于Java中的float和double，分别使用32 和64位来代表单精度浮点数和双精度浮点数

complex64 complex128	-- 复数类型
本例展示了几种类型的变量。 同导入语句一样，变量声明也可以“分组”成一个语法块。

int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
-- int、uint、uintptr 的大小随着操作系统的变化而变化
当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
*/
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	iv     uint       = 0
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	u := iv - 1
	fmt.Print(u)

	fmt.Printf("%T", u)
}
