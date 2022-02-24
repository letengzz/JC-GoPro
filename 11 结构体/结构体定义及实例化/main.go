package main

import "fmt"

//结构体的定义
type person struct {
	name, city string
	age        int8
}

func main() {
	var a person
	a.name = "445"
	a.city = "fbl"
	fmt.Printf("person:%v\n", a)
	fmt.Println(a.name)
	p := person{"dd", "dd", 45}
	fmt.Println(p)
}
