package main

import "fmt"

func main() {
	cityArray := [3]string{"北京", "上海", "广州"}
	//数组的遍历
	//1.for
	for i := 0; i < len(cityArray); i++ {
		fmt.Print(cityArray[i])
	}
	fmt.Println()
	//2.for range
	for index, value := range cityArray {
		fmt.Print(index, value)
	}
}
