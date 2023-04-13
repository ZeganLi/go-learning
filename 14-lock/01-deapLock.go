package main

// 死锁
func main() {

	c := make(chan struct{})
	<-c
}
