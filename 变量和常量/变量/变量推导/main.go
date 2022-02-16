package main

import "fmt"

/*
有时候我们会将变量的类型省略，
这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。
*/
func main() {
	var name = "ff"
	var age = 21
	fmt.Printf("%s %d", name, age)
}
