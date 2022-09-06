package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}
	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE	age = $1", age)
	fmt.Println(rows)
}

var useYear int

func init() {
	//1、 初始化包内变量，保证其后续可用
	useYear = 13

	// 2、在init函数中注册自己的实现的模式降低了Go包对外的直接暴露，尤其是包级变量的暴露，避免了外部通过包级变量对包状态的改动
	//sql.Register("postgres",&Driver{})

	//init函数是一个无参数、无返回值的函数，它的主要目的是保证其所在包在被正式使用之前的初始状态是有效的。
	//一旦init函数在检查包数据初始状态时遇到失败或错误的情况（尽管极少出现），则说明对包的“质检”亮了红灯，如果让包“出厂”，
	//那么只会导致更为严重的影响。
	//因此，在这种情况下，快速失败是最佳选择。我们一般建议直接调用panic或者通过log.Fatal等函数记录异常日志，然后让程序快速退出。
	// 3、 检测错误，快速失败

	//init函数的几个特点：运行时调用、顺序、仅执行一次。
	//Go程序的初始化顺序。
	//init函数是包出厂前的唯一“质检员”。
}
