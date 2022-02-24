package main

import "fmt"

/*
定义方式:

*/
func main() {
	//数组的定义
	var a [5]int  //数组的长度必须是常量，并且长度是数组类型的一部分
	var b [10]int //`[5]int`和`[10]int`是不同的类型
	fmt.Println(b)
	//数组会初始化为int类型的零值
	for i := 0; i < len(a); i++ {
		a[i] = i //赋值
	}
	fmt.Println(a)
	//字符数组定义
	var str [2]string
	str[0] = "world"
	str[1] = "hello"
	fmt.Println(str)
	//第二种方法
	str1 := [2]string{"12", "safas"}
	fmt.Println(str1)
}
