package main

import (
	"io"
	"os"
)

// https://blog.go-zh.org/defer-panic-and-recover
// defer 语句 将函数调用推送到列表中。保存的调用列表在周围函数返回后执行。Defer 通常用于简化执行各种清理操作的函数。

func main() {

}

// CopyFile 打开两个文件并将一个文件的内容复制到另一个文件的函数
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	//如果调用 os.Create 失败，函数将返回而不关闭源文件。
	//这可以通过在第二个 return 语句之前调用 src.Close 来轻松解决，但如果函数更复杂，问题可能不会那么容易被注意到和解决。
	//通过引入 defer 语句，我们可以确保文件始终关闭
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}

func CopyFileDefer(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// 引入defer，来保证文件资源始终能够被关闭，而无论在什么时候返回
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	// 引入defer，来保证文件资源始终能够被关闭，而无论在什么时候返回
	defer dst.Close()

	return io.Copy(dst, src)
}
