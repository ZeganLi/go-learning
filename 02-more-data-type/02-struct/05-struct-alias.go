package main

import (
	"encoding/json"
	"fmt"
)

// Student `` 标记符，因为go中是使用大写字母来控制访问的权限，
// 因此与数据库交互中，或者与其他语言进行json数据通信时，用来转驼峰命名
type Student struct {
	StuName  string `json:"stu_name"`
	StuClass string `json:"stu_class"`
	auth     string `json:"auth"`
}

func main() {
	student := Student{StuName: "张三", StuClass: "三年四班"}
	marshal, _ := json.Marshal(student)
	fmt.Printf("%s", marshal)

}
