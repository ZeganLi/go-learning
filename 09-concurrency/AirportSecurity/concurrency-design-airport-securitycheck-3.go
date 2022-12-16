package main

import "time"

// 排队旅客（passenger）：代表应用的外部请求。
// 机场工作人员：代表计算资源。
// 安检程序：代表应用，必须在获取机场工作人员后才能工作。模拟安检例子中，
//  	   安检程序内部流程包括登机身份检查（idCheck）、人身检查（bodyCheck）和X光机对随身物品的检查（xRayCheck）。
// 安检通道（channel）：每个通道对应一个应用程序的实例。

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTmCost = 180
)

func idCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	println("\tgoroutine-", id, ": idCheck ok")
	return idCheckTmCost
}

func bodyCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	println("\tgoroutine- ", id, "bodyCheck ok")
	return bodyCheckTmCost
}

func xRayCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	println("\tgoroutine- ", id, "xRayCheck ok")
	return xRayCheckTmCost
}

func newAirportSecurityCheckChannel(id string, queue chan struct{}) {

	go func(id string) {
		print("goroutine-", id, ": airportSecurityCheckChannel is ready...\n")
		// 启动x光检查
		queue3, quit3, result3 := start(id, xRayCheck, nil)
		// 启动人身检查
		queue2, quit2, result2 := start(id, bodyCheck, queue3)
		//启动身份检查
		queue1, quit1, result1 := start(id, idCheck, queue2)

		for {
			select {
			case v, ok := <-queue:
				if !ok {
					close(quit1)
					close(quit2)
					close(quit3)

					total := max(<-result1, <-result2, <-result3)
					println("goroutine-", id, ": airportSecurityCheckChannel time cost:", total, "\n")
					print("goroutine-", id, ": airportSecurityCheckChannel closed\n")
					return
				}
				queue1 <- v
			}
		}
	}(id)

}

/**
struct{} 空结构体
使用unsafe.Sizeof(struct{}{}) 来查看空结构体的大小为0
因为空结构体不占据内存空间，因此被广泛作为各种场景下的占位符使用
一是节省资源，二是空结构体本身就具备很强的语义，即这里不需要任何值，仅作为占位符
*/
func start(id string, f func(string) int, next chan<- struct{}) (chan<- struct{}, chan<- struct{}, <-chan int) {
	queue := make(chan struct{}, 10)
	quit := make(chan struct{})
	result := make(chan int)

	go func() {
		total := 0
		for {
			select {
			case <-quit:
				result <- total
				return
			case v := <-queue:
				total += f(id)
				if next != nil {
					next <- v
				}
			}
		}
	}()

	return queue, quit, result
}

func max(args ...int) int {
	n := 0
	for _, v := range args {
		if v > n {
			n = v
		}
	}
	return n
}

func main() {
	passengers := 30
	queue := make(chan struct{}, 30)
	newAirportSecurityCheckChannel("channel1", queue)
	newAirportSecurityCheckChannel("channel2", queue)
	newAirportSecurityCheckChannel("channel3", queue)
	time.Sleep(5 * time.Second)

	for i := 0; i < passengers; i++ {
		queue <- struct{}{}
	}
	time.Sleep(5 * time.Second)
	close(queue)                   // 为了打印各通道的处理时长
	time.Sleep(1000 * time.Second) // 防止main goroutine退出
}
