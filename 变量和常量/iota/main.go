package main

import "fmt"

func main() {
	const (
		n1 = iota
		n2
		n3
	)
	fmt.Println(n1, n2, n3)
	//	使用_跳过某些值
	const (
		m1 = iota
		m2
		_
		m4
	)
	fmt.Println(m1, m2, m4)
	//	iota声明中间插队
	const (
		k1 = iota //0
		k2 = 100  //100
		k3 = iota //2
		k4        //3
	)
	const k5 = iota
	fmt.Println(k1, k2, k3, k4, k5)
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println(KB, MB, GB, TB, PB)
	//多个iota定义在一行
	const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
	fmt.Println(a, b, c, d, e, f)
}
