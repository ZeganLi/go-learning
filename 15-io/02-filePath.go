package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	filePath1 := "15-io/files/美女.jpg"
	filePath2 := "D:\\go-project\\go-tour\\15-io\\files\\美女.jpg"

	// 是否是绝对路径
	fmt.Println(filepath.IsAbs(filePath1))
	fmt.Println(filepath.IsAbs(filePath2))

	// 获取相对路径
	fmt.Println(filepath.Rel("D:\\", filePath2))

	// 获取绝对路径
	fmt.Println(filepath.Abs(filePath1))
	fmt.Println(filepath.Abs(filePath2))

	// 拼接路径
	fmt.Println(filepath.Join(filePath1, "."))
	fmt.Println(filepath.Join(filePath2, ".."))
	fmt.Println(filepath.Join(filePath2, "."))
	fmt.Println(filepath.Join("E:\\", filePath1))
}
