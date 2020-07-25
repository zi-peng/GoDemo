package main

import "fmt"

/*返回最大值*/
func max(num1, num2 int) int {

	/* 定义局部变量 */
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
func swap(y, x string) (string, string) {
	return y, x
}
func main() {
	//定义局部变量
	a := 100
	b := 200
	var ret int
	ret = max(a, b)
	c, d := swap("Google", "Runoob")
	fmt.Println(c, d)
	fmt.Printf("最大值是:%d\n", ret)
}
