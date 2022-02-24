package main

import "fmt"

func main() {
	//光声明map类型 但是没有初始化 a就是初始值nil
	var a map[string]int
	//没有初始化不能直接添加键值对
	//a["s"]=100
	fmt.Println(a == nil)
	//map的初始化
	b := make(map[int]bool, 5)
	fmt.Println(b == nil)
	//map中添加键值对
	b[1] = true
	b[2] = false
	//值的Go语法表示
	fmt.Printf("%#v\n", b)
	fmt.Printf("type:%T\n", b)
	//声明map的同时完成初始化
	c := map[string]int{
		"第一个": 1,
		"第二个": 2, //}
		//若后大括号在下面则添加一个`,`否则要跟在值的后面,如上表示
	}
	fmt.Printf("%#v\n", c)
	fmt.Printf("type:%T\n", c)
}
