package main

import "fmt"

func main() {

	numbers := [256]int{'a': 8, 'b': 7, 'c': 4, 'd': 3, 'e': 2, 'y': 1, 'x': 5}

	// [10]float{-1, 0, 0, 0, -0.1, -0.1, 0, 0.1, 0, -1}
	// 因为最后一个值的下标为9，所以推导出数组的元素个数为10
	fNumbers := [...]float32{-1, 4: -0.1, -0.1, 7: 0.1, 9: 3.14}

	// $GOROOT/src/sort/search_test.go
	var data = []int{0: -10, 1: -5, 2: 0, 3: 1, 4: 2, 5: 3, 6: 5, 7: 7, 8: 11, 9: 100, 10: 100, 11: 100, 12: 1000, 13: 10000}

	var sdata = []string{0: "f", 1: "foo", 2: "foobar", 3: "x"}

	fmt.Printf("numbers = %v,长度为 = %d \n", numbers, len(numbers))
	fmt.Printf("fnumbers = %v,长度为 = %d \n", fNumbers, len(fNumbers))
	fmt.Printf("data = %v,长度为 = %d \n", data, len(data))
	fmt.Printf("sdata = %v,长度为 = %d \n", sdata, len(sdata))

	fmt.Println('a')
}
