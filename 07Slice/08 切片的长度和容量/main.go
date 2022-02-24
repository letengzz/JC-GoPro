package main

import "fmt"

func main() {
	var a = [5]int{12, 4, 67, 89, 66}
	b := a[0:2]
	c := make([]int, 1, 14)

	fmt.Printf("%T %T\n", b, c)
	//通过len()函数获取切片的长度
	fmt.Println("长度:", len(b))
	fmt.Println("长度:", len(c))
	//通过cap()函数获取切片的容量
	fmt.Println("容量:", cap(b))
	fmt.Println("容量:", cap(c))
}
