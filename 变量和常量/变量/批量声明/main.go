package main

import "fmt"

/*Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且Go语言的变量声明后必须使用*/
func main() {
	var (
		a string
		b int
		c bool
	)
	fmt.Printf("%s %d %t",a,b,c)
}
