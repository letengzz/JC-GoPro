package main

import "fmt"

func main() {
	a := map[string]int{
		"ff": 100,
		"aa": 200,
	}
	fmt.Println(a)
	//删除键值对
	delete(a, "aa")
	fmt.Println(a)
}
