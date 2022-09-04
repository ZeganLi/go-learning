package main

import (
	"fmt"
)

func main() {

	a()        //0
	b()        // 3210
	print(c()) //2

	//defer的 其他用途（除了前面给出的 file.Close 示例）包括释放互斥锁：
	//mu.Lock()
	//defer mu.Unlock()
}

// 1.当 defer 语句被执行时，延迟函数的参数被计算。
// 在这个例子中，当 Println 调用被延迟时，表达式“i”被计算,因此放入栈中的数据为0，而不是i++之后的值。
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 2.延迟函数调用在周围函数返回后按后进先出的顺序执行。
// defer栈，后入先出
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// 3.延迟函数可以读取并分配给返回函数的命名返回值。
// 在此示例中，延迟函数在周围函数返回后递增返回值 i。 因此，此函数返回 2：
// 这样方便修改函数的错误返回值
func c() (i int) {
	// defer 函数可以在最后对即将要返回的值进行操作
	defer func() { i++ }() // 2
	defer fmt.Println(i)   // 0
	return 1
}
