package main

import "fmt"

/*
Go语言的变量声明格式为：
	var 变量名 变量类型
Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且Go语言的变量声明后必须使用
*/
func main() {
	var name string = "tom"
	var age int = 15
	var isOk bool = false
	fmt.Printf("我的名字%s age:%d Ok:%t", name, age,isOk)
}
