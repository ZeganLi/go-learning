package main

import "fmt"

// MyCombinationError
// 在一些场景下，错误处理者需要从错误值中提取出更多信息以帮助其选择错误处理路径，这时他们可以自定义错误类型来满足需求.
// 通过组合error接口的方式，来扩展错误的信息
// error接口是 错误值提供者 与 错误值检视者 之间的契约。
// error接口的实现者负责提供错误上下文供负责错误处理的代码使用。
// 这种错误上下文与error接口类型的分离体现了Go设计哲学中的“正交”理念。
type MyCombinationError struct {
	Name  string
	City  string
	Value string
	Err   error
}

func main() {

	ce := MyCombinationError{}
	ce.Name = "张三"
	ce.City = "北京"
	ce.Value = "滑冰"
	ce.Err = fmt.Errorf("%s在%s因%s死了", ce.Name, ce.City, ce.Value)

	fmt.Println(ce)
}
