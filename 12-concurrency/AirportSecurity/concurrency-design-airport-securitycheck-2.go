package main

//
//import "time"
//
//// 排队旅客（passenger）：代表应用的外部请求。
//// 机场工作人员：代表计算资源。
//// 安检程序：代表应用，必须在获取机场工作人员后才能工作。模拟安检例子中，
////  	   安检程序内部流程包括登机身份检查（idCheck）、人身检查（bodyCheck）和X光机对随身物品的检查（xRayCheck）。
//// 安检通道（channel）：每个通道对应一个应用程序的实例。
//
//const (
//	idCheckTmCost   = 60
//	bodyCheckTmCost = 120
//	xRayCheckTmCost = 180
//)
//
//func idCheck(id int) int {
//	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
//	println("\tgoroutine-", id, ": idCheck ok")
//	return idCheckTmCost
//}
//
//func bodyCheck(id int) int {
//	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
//	println("\tgoroutine- ", id, "bodyCheck ok")
//	return bodyCheckTmCost
//}
//
//func xRayCheck(id int) int {
//	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
//	println("\tgoroutine- ", id, "xRayCheck ok")
//	return xRayCheckTmCost
//}
//
//func airportSecurityCheck(id int) int {
//	println("airportSecurityCheck ...")
//	total := 0
//	total += idCheck(id)
//	total += bodyCheck(id)
//	total += xRayCheck(id)
//
//	println("airportSecurityCheck OK")
//	return total
//}
//
///**
//struct{} 空结构体
//使用unsafe.Sizeof(struct{}{}) 来查看空结构体的大小为0
//因为空结构体不占据内存空间，因此被广泛作为各种场景下的占位符使用
//一是节省资源，二是空结构体本身就具备很强的语义，即这里不需要任何值，仅作为占位符
//*/
//func start(id int, f func(int) int, queue <-chan struct{}) <-chan int {
//	c := make(chan int)
//	go func() {
//		total := 0
//		for {
//			_, ok := <-queue
//			if !ok {
//				c <- total
//				return
//			}
//
//			total += f(id)
//		}
//	}()
//
//	return c
//}
//
//func max(args ...int) int {
//	n := 0
//	for _, v := range args {
//		if v > n {
//			n = v
//		}
//	}
//	return n
//}
//
//func main() {
//	total := 0
//	passengers := 30
//	c := make(chan struct{})
//	c1 := start(1, airportSecurityCheck, c)
//	c2 := start(2, airportSecurityCheck, c)
//	c3 := start(3, airportSecurityCheck, c)
//	for i := 0; i < passengers; i++ {
//		c <- struct{}{}
//	}
//
//	close(c)
//
//	total = max(<-c1, <-c2, <-c3)
//	println("total time cost:", total)
//}
