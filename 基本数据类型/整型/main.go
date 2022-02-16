package main

import "fmt"

func main() {
	//无符号类型
	var ua uint8 = 255
	var ub uint16 = 256
	var uc uint32 = 666
	var ud uint64 = 777
	fmt.Printf("uint8:%d uint16:%d uint32:%d uint64:%d\n", ua, ub, uc, ud)
	//有符号类型
	var ia int8 = 127
	var ib int16 = 128
	var ic int32 = 666
	var id int64 = 777
	fmt.Printf("int8:%d int16:%d int32:%d int64:%d\n", ia, ib, ic, id)
	//特殊整型 不同平台上的差异
	var a uint = 123
	var b int = 456
	//var c uintptr =
	fmt.Printf("uint:%d int:%d\n", a, b)
	//数字字面量语法
	//十进制
	q := 10
	fmt.Printf("%d\n", q)
	//八进制 以0开头
	w := 075
	e := 0o52
	fmt.Printf("%o %o\n", w, e)
	//十六进制 以0x开头
	r := 0xf0
	fmt.Printf("%X\n", r)
	//特殊
	t := 123_456
	y := 0x1p-2
	fmt.Printf("%d\n", t)
	fmt.Printf("%f", y)
}
