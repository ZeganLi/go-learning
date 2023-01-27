package main

import (
	"errors"
	"fmt"
)

// 错误处理方需要错误值提供更多的错误上下文，上面"哨兵"错误处理策略和错误值构造方式将无法满足

// 我们需要通过自定义错误类型的构造错误值的方式来提供更多的错误上下文信息，并且由于错误值均通过error接口变量统一呈现，要得到底层错误类型携带的错误上下文信息，
// 错误处理方需要使用Go提供的类型断言机制（type assertion）或类型选择机制（typeswitch） 这种错误处理笔者称之为错误值类型检视策略
func main() {

	var err = &MyTypeAssertionError{"my error type"}
	err1 := fmt.Errorf("wrap err1: %w", err)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	var e *MyTypeAssertionError
	// errors.As函数沿着err2所在错误链向上找到了被包装到最深处的错误值，并将err2与其类型*MyTypeAssertionError成功匹配
	// 如果你使用的是Go 1.13及后续版本,请尽量使用errors.As方法去检视某个错误值是不是某个自定义错误类型的实例

	// errors.Is是检视某个<<错误值>>是不是某个特定"哨兵"的<<错误值>>
	// errors.As是检视某个<<错误类型>>是不是某个<<自定义错误类型的实例>>
	// 注意二者的区别；Is是错误值之间的比较，As是错误类型之间的比较
	if errors.As(err2, &e) {
		println("MyError is on the chain of err2 ")
		println(e == err)
		return
	}
	println("MyError is not on the chain of err2 ")
}

// MyTypeAssertionError 一般自定义导出的错误类型以XXXError的形式命名
type MyTypeAssertionError struct {
	e string
}

func (e *MyTypeAssertionError) Error() string {
	return e.e
}
