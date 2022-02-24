package main

import "fmt"

func main() {
	//判断某个键存不存在
	a := make(map[string]int, 5)
	a["吴亦凡"] = 100
	a["gg"] = 50
	//判断在不在
	value, ok := a["嘿嘿"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("嘿嘿在a中", value)
	} else {
		fmt.Println("没有")
	}
}
