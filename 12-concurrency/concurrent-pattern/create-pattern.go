package main

type T struct {
}

// 创建模式
func main() {
	c := spawn(func() {})
	// 使用c来通信
	<-c
}

func spawn(f func()) chan struct{} {
	c := make(chan struct{})
	go f()
	return c
}
