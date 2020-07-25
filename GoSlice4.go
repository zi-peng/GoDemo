package main

import "fmt"

//=====切片 append() 和 copy() 函数
//1.append()  //这里的追加指的是具体的值，使用是  append(numbers,1) numbers 表示的是要追加的切片，1代表的是追加的值
//2.copy()   //copy的使用方法是 copy(numbers,numbers1) numbers1是数据源，number是接收的值的一方
func main() {
	var numbers []int
	printSlice5(numbers)
	//1.允许追加空切片
	// numbers =append(numbers,0)
	printSlice5(numbers)
	//2.向切片添加一个元素
	numbers = append(numbers, 1)
	printSlice5(numbers)
	//3. 向切片添加多个元素
	numbers = append(numbers, 2, 3, 4)
	printSlice5(numbers)
	//创建切片 numbers1 是之前切片的两倍容量, 扩容比实际存储多一倍
	numbers1 := make([]int, len(numbers), cap(numbers)*2)
	//拷贝numbers的内容到numbers1
	copy(numbers1, numbers) //目标文件,源文件
	printSlice5(numbers1)
}
func printSlice5(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
