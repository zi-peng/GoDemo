package main

import "fmt"

//=======空切片示例========
//1.如果定义的是空的切片，那么对应的实际长度和最大上限值也都会是0
//2.切片显示的是一个空的集合 []
func main() {
	var numbers []int
	printSlice1(numbers)
	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

func printSlice1(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
