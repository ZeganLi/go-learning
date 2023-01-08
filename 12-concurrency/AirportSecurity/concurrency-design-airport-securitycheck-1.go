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
//func idCheck() int {
//	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
//	println("\tidCheck ok")
//	return idCheckTmCost
//}
//
//func bodyCheck() int {
//	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
//	println("\tbodyCheck ok")
//	return bodyCheckTmCost
//}
//
//func xRayCheck() int {
//	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
//	println("\txRayCheck ok")
//	return xRayCheckTmCost
//}
//
//func airportSecurityCheck() int {
//	println("airportSecurityCheck ...")
//	total := 0
//	total += idCheck()
//	total += bodyCheck()
//	total += xRayCheck()
//
//	println("airportSecurityCheck OK")
//	return total
//}
//
//func main() {
//	total := 0
//	passengers := 30
//
//	for i := 0; i < passengers; i++ {
//		total += airportSecurityCheck()
//	}
//
//	println("total time cost:", total)
//}
