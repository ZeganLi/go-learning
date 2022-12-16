package main

import (
	"time"
)

//go f(x, y, z) 会启动一个新的Go程并执行 f(x,y,z)
// x y z 的求值发生在当前的Go程中，而f的执行发生在新的Go程中

// Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。sync 包提供了这种能力
func main() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		println(s)
	}
}
