package main

import (
	"fmt"
	"strings"
)

//定义一个函数它的返回值是一个函数
func a(name string) func() {
	//name := "hh"
	return func() {
		fmt.Println("HEllO", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	//闭包 = 函数 + 外层变量的引用
	r := a("hh") //r此时就是一个闭包
	r()          //相当于执行了a函数内部的匿名函数
	r1 := makeSuffixFunc(".txt")
	ret := r1("nih")
	fmt.Println(ret)

	rc1, rc2 := calc(200)
	ret1 := rc1(200) //base = 200 + 200
	fmt.Println(ret1)
	ret2 := rc2(600) //base = 400 - 600
	fmt.Println(ret2)

}
