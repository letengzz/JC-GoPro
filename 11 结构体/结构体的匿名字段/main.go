package main

import "fmt"

//结构体的匿名字段
type Person struct {
	string
	int
	//string
	//结构体匿名字段不能重复
}

func main() {
	p1 := Person{
		string: "小王子",
		int:    18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
}
