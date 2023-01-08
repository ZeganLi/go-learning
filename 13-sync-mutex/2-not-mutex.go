package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	var count int
	for i := 0; i < 10000; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			count++
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(count)

}
