package main

import "fmt"

/*方法是一种作用于特定类型的函数
非本地类型不能定义方法

*/

// Person 是一个结构体
type Person struct {
	name string
	age  int8
}

//NewPerson
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 是为Person类型定义方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是运动员！", p.name)
}

// Setage 是一个修改年龄的方法
//指针接收者指的是接收者的类型是指针
func (p *Person) Setage(newAge int8) {
	p.age = newAge
}

// Setage 使用值接收者来修改年龄的方法
func (p Person) Setage2(newAge int8) {
	p.age = newAge
}

//为任意类型添加方法
type myInt int

func (m myInt) sayHi() {
	fmt.Println("hi")
}
func main() {
	p1 := NewPerson("bingbing", int8(20))
	//(*p1).Dream()
	p1.Dream()

	//调用修改年龄的方法
	fmt.Println(p1.age)
	p1.Setage(int8(10))
	fmt.Println(p1.age)
	p1.Setage2(int8(30))
	fmt.Println(p1.age)

	var m1 myInt
	m1 = 100
	m1.sayHi()
}
