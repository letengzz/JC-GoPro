package main

import "fmt"

//定义一个函数它的返回值是一个函数
func a() func() {
	return func() {
		fmt.Println("HEllO")
	}
}
func main() {
	r := a()
	r() //相当于执行了a函数内部的匿名函数
}
