package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var m sync.Mutex
var c chan struct{}

func main() {
	//go lockAdd(&count)
	//go lockSub(&count)
	for i := 0; i < 10; i++ {
		go chanAdd(&count)
		go chanSub(&count)

		time.Sleep(500 * time.Millisecond)
		fmt.Println("count:", count)
	}
}

func chanSub(count *int) {
	func() {
		for i := 0; i < 100000; i++ {
			<-c
			*count--
		}
	}()
}

func chanAdd(count *int) {
	func() {
		for i := 0; i < 100000; i++ {
			c <- struct{}{}
			*count++
		}
	}()
}

func lockSub(count *int) {
	func() {
		for i := 0; i < 100000; i++ {
			m.Lock()
			*count--
			m.Unlock()
		}
	}()
}

func lockAdd(count *int) {
	func() {
		for i := 0; i < 100000; i++ {
			m.Lock()
			*count++
			m.Unlock()
		}
	}()
}
