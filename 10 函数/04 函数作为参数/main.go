package main

import "fmt"

func testGlobal() {
	num := 200
	name := "刘德华"
	fmt.Println("testGlobal中的变量是", num)
	fmt.Println(name)
}
func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func main() {
	//函数是可以作为变量的
	abc := testGlobal
	fmt.Printf("%T\n", abc)
	abc()

	r1 := calc(100, 200, add)
	r2 := calc(200, 14, sub)
	fmt.Println(r1)
	fmt.Println(r2)
}
