package main

import "fmt"

func main() {
	const p = 3.14
	//多个常量也可以一起声明
	const (
		e  = 50
		f  = 60
		n1 = 70
		n2
		n3
		//`const`同时声明多个常量时，如果省略了值则表示和上面一行的值相同。
	)
	fmt.Println(p, e, f, n1, n2, n3)
}
