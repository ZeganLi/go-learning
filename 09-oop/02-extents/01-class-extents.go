package main

import "fmt"

type Human struct {
	Name string
	age  uint8
}

type SupperMan struct {
	Human
	sex byte
}

func main() {

	//human := Human{"zhang3", 18}
	//fmt.Printf("Human %T,Name:%v,Age:%d", human, human.Name, human.age)

	//supperMan := SupperMan{Human{"李四", 20}}
	//fmt.Printf("Human %T,Name:%v,Age:%d", supperMan, supperMan.Name, supperMan.age)

	var supperMan SupperMan
	supperMan.Name = "王五"
	supperMan.age = 88
	supperMan.sex = 1
	fmt.Printf("Human %T,Name:%v,Age:%d,sex:%b", supperMan, supperMan.Name, supperMan.age, supperMan.sex)

}
