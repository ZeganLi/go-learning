package pkg1

import "fmt"

var (
	_  = constInitCheck() // 不会输出，相当于没有引用
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

const (
	P1 = "c1"
	c2 = "c2"
)

func constInitCheck() string {
	if P1 != "" {
		fmt.Println("main: const c1 init")
	}
	if P1 != "" {
		fmt.Println("main: const c2 init")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("main: var %s init\n", name)
	return name
}

func main() {

	fmt.Println("main:pkg1")
}

func init() {
	fmt.Println("init:pkg1")
}
