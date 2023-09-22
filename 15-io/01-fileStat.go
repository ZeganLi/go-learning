package main

import (
	"fmt"
	"os"
)

func main() {

	filePath := "15-io/files/美女.jpg"

	printMessage(filePath)
}

func printMessage(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("数据类型是：%T \n", fileInfo)
		fmt.Println("文件名：", fileInfo.Name())
		fmt.Println("是否是目录：", fileInfo.IsDir())
		fmt.Println("文件大小：", fileInfo.Size())
		fmt.Println("文件权限：", fileInfo.Mode())
		fmt.Println("文件最后修改时间：", fileInfo.ModTime())
	}
}
