package main

import "fmt"

//结构体的定义
type person struct {
	name, city string
	age        int8
}

func main() {
	//键值对初始化
	p1 := person{
		name: "nihao",
		city: "moa",
		age:  10,
	}
	p2 := &person{
		name: "nihao",
		city: "moa",
		age:  10,
	}
	fmt.Printf("%#v %#v\n", p1, p2)
	fmt.Println(p1.name, p1.city, p1.age)
	//使用值的列表进行初始化
	p3 := person{
		"nihao",
		"dd",
		50,
	}
	fmt.Printf("%#v\n", p3)
}
