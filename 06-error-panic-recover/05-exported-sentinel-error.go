package main

import (
	"errors"
	"fmt"
)

// 如果需要对错误值进行查看，并选择根据不同的错误执行不同的业务逻辑时，就需要详细的查看返回的错误值
func main() {

	// 对错误值进行检视
	_, err := peek(0)
	if err != nil {
		switch err.Error() {
		case "i 小于 0":
			fmt.Println(err)
		case "i 等于 0 ": // 当错误构造方修改了错误的内容,造成了错误处理方行为发生了变化
			fmt.Println(err)
		case "i 大于 0":
			fmt.Println(err)
		default:
			fmt.Println("默认值")
		}
	}

	// 上述 demo 存在很重要的问题，错误值构造方不经意间的一次错误描述字符串的改动，都会造成错误处理方处理行为的变化
	// 并且这种通过字符串比较的方式对错误值进行检视的性能也很差
	// -----------------------------------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------------------
	// 	Go标准库中采用了定义导出的"哨兵"错误值的方式来对上述问题进行解决
	// 使用哨兵模式对错误值进行检视
	_, err = sentinel(1)
	if err != nil {
		switch err {
		case ErrLtZero:
			fmt.Println(err)
		case ErrEqZero:
			fmt.Println(err)
		case ErrRtZero:
			fmt.Println(err)
		default:
			fmt.Println("默认值")
		}
	}

	// -----------------------------------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------------------
	// 	从Go 1.13版本开始，标准库errors包提供了Is方法用于错误处理方对错误值进行检视。
	// 	Is方法类似于将一个error类型变量与“哨兵”错误值的比较
	_, err = sentinel(1)
	if errors.Is(err, ErrRtZero) {
		fmt.Println(err)
	}

	if errors.Is(err, ErrEqZero) {
		fmt.Println(err)
	}

	if errors.Is(err, ErrRtZero) {
		fmt.Println(err)
	}

	fmt.Println("默认值")

	// -----------------------------------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------------------
	// 	如果error类型变量的底层错误值是一个包装错误（wrap error），errors.Is方法会沿着该包装错误所在错误链（errorchain）
	// 	与链上所有被包装的错误（wrapped error）进行比较，直至找到一个匹配的错误
	err1 := fmt.Errorf("wrap err1: %w", ErrSentinel)
	err2 := fmt.Errorf("wrap err2: %w", err1)
	if errors.Is(err2, ErrSentinel) {
		println("err is ErrSentinel") // 输出 "err is ErrSentinel" 说明找到了最深处的错误值"ErrSentinel"
		return
	}
	println("err is not ErrSentinel")
	// 如果你使用的是Go1.13及后续版本，请尽量使用errors.Is方法检视某个错误值是不是某个特定的“哨兵”错误值

	// -----------------------------------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------------------
	// -----------------------------------------------------------------------------------------------------------------
}

var ErrSentinel = errors.New("the underlying sentinel error")

// “哨兵”错误值变量以ErrXXX格式命名
var (
	ErrLtZero = errors.New("i 小于 0")
	ErrEqZero = errors.New("i 等于 0")
	ErrRtZero = errors.New("i 大于 0")
)

func peek(i int) (int, error) {

	if i < 0 {
		return i, errors.New("i 小于 0")
	}

	if i == 0 {
		return i, errors.New("i 等于 0")
	}

	if i > 0 {
		return i, errors.New("i 大于 0")
	}

	return i, nil
}

func sentinel(i int) (int, error) {
	if i < 0 {
		return i, ErrLtZero
	}

	if i == 0 {
		return i, ErrEqZero
	}

	if i > 0 {
		return i, ErrRtZero
	}

	return i, nil
}
