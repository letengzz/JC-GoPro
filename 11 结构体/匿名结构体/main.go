package main

import "fmt"

func main() {
	//匿名结构体
	var user struct {
		name    string
		married bool
	}
	user.name = "你好好"
	user.married = false
	fmt.Println(user)
}
