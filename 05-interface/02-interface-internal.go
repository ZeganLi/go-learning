package main

import "fmt"

func main() {
	//printNilInterface()
	//printEmptyInterface()
	//printNonEmptyInterface()
	//printEmptyInterfaceAndNonEmptyInterface()
	var i interface{}
	zs := T{n: 18, s: "张三"}
	ls := T{n: 18, s: "李四"}
	i = zs
	println("zs: ", i)
	i = ls
	println("ls: ", i)
	println(zs == ls)
}

/**
nil接口变量
*/
func printNilInterface() {
	// nil接口变量
	var i interface{} // 空接口类型
	var err error     // 非空接口类型
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i = err:", i == err)
	println("")

	// 无论是空接口类型变量还是非空接口类型变量，一旦变量值为nil，那么它们内部表示均为(0x0,0x0)，
	// 即类型信息和数据信息均为空。因此上面的变量i和err等值判断为true
}

/**
空接口类型变量 对于空接口类型变量，只有在_type和data所指数据内容一致（注意：不是数据指针的值一致）的情况下，两个空接口类型变量之间才能画等号
*/
func printEmptyInterface() {
	var eif1 interface{} // 空接口类型
	var eif2 interface{} // 空接口类型
	var n, m int = 17, 18
	eif1 = n
	eif2 = m
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2)
	eif2 = 17
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2)

	eif2 = int64(17)
	println("eif1:", eif1)
	println("eif2:", eif2)
	// 二者类型不一致，因此是false
	println("eif1 = eif2:", eif1 == eif2)
	println("")
}

/**
非空接口类型变量 与空接口类型变量一样，只有在tab和data所指数据内容一致的情况下，两个非空接口类型变量之间才能画等号
*/
func printNonEmptyInterface() {
	var err1 error // 非空接口类型
	var err2 error // 非空接口类型
	println(err1 == nil)
	println(err2 == nil)

	//针对这种赋值，println输出的err1是(0x10ed120, 0x0)，
	//即非空接口类型变量的类型信息并不为空，数据指针为空，因此它与nil(0x0,0x0)之间不能画等号
	err1 = (*T)(nil)
	println("err1:", err1)
	println("err1 = nil:", err1 == nil)
	//err1 = T(5)
	//err2 = T(6)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)
	err2 = fmt.Errorf("%d\n", 5)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)
	println("")
}

/**
空接口类型变量与非空接口类型变量的等值比较
空接口类型变量和非空接口类型变量内部表示的结构有所不同（第一个字段：_type vs. tab），似乎一定不能相等。
但Go在进行等值比较时，类型比较使用的是eface的_type和iface的tab._type，
因此就像我们在这个例子中看到的那样，当eif和err都被赋值为T(5)时，两者之间是可以画等号的。
*/
func printEmptyInterfaceAndNonEmptyInterface() {
	//var eif interface{} = T(5)
	//var err error = T(5)
	//println("eif:", eif)
	//println("err:", err)
	//println("eif = err:", eif == err)
	//err = T(6)
	//println("eif:", eif)
	//println("err:", err)
	//println("eif = err:", eif == err)
}

type T struct {
	n int
	s string
}

func (T) M1() {
}

func (T) M2() {
}

func (T) Error() string {
	return ""
}

type NonEmptyInterface interface {
	M1()
	M2()
	error
}
