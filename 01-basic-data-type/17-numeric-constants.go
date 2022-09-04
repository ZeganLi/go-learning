package main

import "fmt"

/*
数值常量是高精度的 值。

一个未指定类型的常量由上下文来决定其类型。 -- 使用的时候再确定常量的类型

再尝试一下输出 needInt(Big) 吧。

（int 类型最大可以存储一个 64 位的整数，有时会更小。）

（int 可以存放最大64位的整数，根据平台不同有时会更少。）
*/
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = 1 >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println(needInt(Small))
	fmt.Printf("needInt返回的数据类型：%T \n", needInt(Small)) // int
	fmt.Println(needFloat(Small))
	fmt.Printf("needFloat返回的数据类型：%T", needFloat(Small)) // float64
	fmt.Println(needFloat(Big))
	fmt.Printf("needFloat返回的数据类型：%T", needFloat(Big)) // float64

	// 无类型常量，个人感觉如Pi = 3.1415926这个字面量而言，可以根据使用的场景转换为具体的类型，而无需像Java那样需要显示的进行类型转换(toString)。
	// 如：
	// 2r/Π  此时pi应为float
	// "pi的值为：Π"  此时pi的类型应为字符类型

	//无类型常量是Go语言推荐的实践，它拥有和字面值一样的灵活特性，可以直接用于更多的表达式而不需要进行显式类型转换，从而简化了代码编写。
	//此外，按照Go官方语言规范的描述，数值型无类型常量可以提供比基础类型更高精度的算术运算，至少有256 bit的运算精度。
}
