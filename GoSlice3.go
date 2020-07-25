package main

import "fmt"

//=======切片截取=========
//可以通过设置下限及上限设置来截取切片【lower-bound(下界):upper-bound(上界)】
//扩展(引用一个高中数学概念)：
//满足不等式a<X<b的所有实数X所组成的集合叫开区间，用记号（a，b）表示； 开区间取包括下界数，不包括上界数
//满足不等式a≤X≤b的所有实数X所组成的集合叫闭区间，用记号[a,b] 表示； 闭区间包括条件本身
func main() {
	//创建切片  []表示是切片类型  int 切片所定义的数据类型
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	printSlice2(numbers)
	//1.打印未截取的切片
	fmt.Println("numbers ==", numbers)
	//2.打印子切片从索引1(包含) 到索引4(不包含)
	fmt.Println("numbers[1:4]==", numbers[1:4])
	//3.默认下限为0,相当于是从[0,3]
	fmt.Println("numbers[:3]==", numbers[:3])
	//4.默认上限为Len(s)
	fmt.Println("numbers[4:]==", numbers[4:])
	//====使用make创建切片 len 0 , cap 5 ====
	numbers1 := make([]int, 0, 5)
	printSlice2(numbers1)
	//1.打印子切片从索引0(包含) 到索引 2(不包含)
	numbers2 := numbers[:2]
	printSlice2(numbers2)
	//2.打印子切片从索引 2(包含) 到索引5(不包含)
	numbers3 := numbers[2:5]
	printSlice2(numbers3)

}

// %d 输出整数  %v 输出实际类型的值
func printSlice2(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
