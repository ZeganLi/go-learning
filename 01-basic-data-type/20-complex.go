package main

import "fmt"

// go中还有复数的数据类型：complex64 和complex128
// 默认的复数类型为complex128
// 复数的更多概念 参考百度
func main() {

	// 声明复数
	//var c1 complex128

	// 直接赋值
	//c2 := complex(0.123, 1.321)

	// 声明并赋值
	//var c3 complex128 = complex(1, 2)

	// 复数由实部(r)和虚部(i)两部分构成
	// 通过内置函数real获取复数的实部部分
	//r := real(c2)
	// 通过内置函数imag获取复数的虚部部分
	//i := imag(c2)

	// 复数的运算
	var x complex128 = complex(1, 2) // 1+2i  (1+2i)为复数表达式，1 实部 2 虚部
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"
	fmt.Println(imag(x * y))         // "10"

	fmt.Println(complex(real(x*y), imag(x*y)))

	//复数也可以用==和!=进行相等比较，只有两个复数的实部和虚部都相等的时候它们才是相等的。

	// math/cmplx包中 go语言内置对复数操作的公共方法
}
