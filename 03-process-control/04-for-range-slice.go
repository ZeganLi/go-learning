package main

import (
	"fmt"
)

func main() {

	var sa = make([]int, 10, 10)
	sa = append(sa, 1, 2, 3, 4, 5, 6)

	sb := []string{"haha", "huhu", "dff", "dfdf"}

	for i, s := range sb {
		fmt.Println(i, s)
	}
}
