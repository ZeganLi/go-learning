package main

// 向切片追加切片
func main() {

	var slice1 []int                      // 定义一个切片 其指向nil
	var appArr = [2]int{1, 2}             //定义一个int数组
	slice1 = append(slice1, appArr[:]...) // 向slice1切片中添加基于appArr的切片，需要在后面使用"...",来展开为参数列表

	a := []string{"John", "Paul"}
	b := []string{"George", "Ringo", "Pete"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	// a == []string{"John", "Paul", "George", "Ringo", "Pete"}
}

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v) // 由于切片的零值 nil 用起来就像一个长度为零的切片，我们可以声明一个切片变量然后在循环 中向它追加数据：

		}
	}
	return p
}
