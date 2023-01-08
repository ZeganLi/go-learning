package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t.Left != nil {
			walker(t.Left)
		}
		ch <- t.Value
		if t.Right != nil {
			walker(t.Right)
		}
	}
	walker(t)
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}
func main() {
	fmt.Println("1 and 1 same: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("1 and 2 same: ", Same(tree.New(1), tree.New(2)))

}
