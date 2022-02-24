package main

import "fmt"

type test struct {
	a, b, c int8
}

func main() {
	n := test{
		a: 12,
		b: 13,
		c: 20,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
}
