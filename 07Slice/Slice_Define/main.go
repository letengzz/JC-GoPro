package main

import "fmt"

func main() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{1, 2, 3, 4}   //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
	//基于数组得到切片
	var intArr = [5]int{1, 2, 3, 4}
	intSlice := intArr[1:3]
	fmt.Println(intSlice)
	fmt.Printf("%T\n", intSlice)
	//切片再次切片
	intSlice2 := intSlice[0:len(intSlice)]
	fmt.Println(intSlice2)
	fmt.Printf("%T\n", intSlice2)
	//通过make函数构造切片
	d := make([]int, 3, 10)
	fmt.Println(d)
	fmt.Printf("%T\n", d)

}
