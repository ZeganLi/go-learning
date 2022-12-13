package main

import "fmt"

/**
在调用方法的时候，传递的接收者本质上都是副本，只不过一个是这个值副本，一是指向这个值指针的副本。
指针具有指向原有值的特性，所以修改了指针指向的值，也就修改了原有的值。
我们可以简单的理解为值接收者使用的是值的副本来调用方法，而指针接收者使用实际的值来调用方法。
*/
func main() {
	p := person1{name: "张三"}
	p.modify() //指针接收者，修改有效
	fmt.Println(p.String())
}

type person1 struct {
	name string
}

func (p person1) String() string {
	return "the person name is " + p.name
}

func (p *person1) modify() {
	p.name = "李四"
}
