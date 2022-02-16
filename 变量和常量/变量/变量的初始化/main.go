package main

import "fmt"

/*
1. 整型和浮点型变量的默认值为`0`。
2. 字符串变量的默认值为`空字符串`。
3. 布尔型变量默认为`false`。
4. 切片、函数、指针变量的默认为`nil`。
在声明变量的时候为其指定初始值。变量初始化的标准格式如下：
	var 变量名 类型 = 表达式
*/
func main() {
	var name string = "nihao"
	var age int = 50
	fmt.Printf("%s %d\n", name, age)
	//一次初始化多个变量
	var a, b = "aad", 20
	fmt.Printf("%s,%d", a, b)
}
