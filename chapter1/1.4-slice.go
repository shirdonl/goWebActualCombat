//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

import "fmt"

func main() {
	//var a  = [3]int{1, 2, 3}
	//fmt.Println(a, a[1:2])

	//var sliceBuilder [20]int
	//for i := 0; i < 20; i++ {
	//	sliceBuilder[i] = i + 1
	//}
	//fmt.Println(sliceBuilder[5:15]) // 区间
	//fmt.Println(sliceBuilder[15:]) // 中间到尾部的所有元素
	//fmt.Println(sliceBuilder[:2]) // 开头到中间指定位置的所有元素

	//a := []int{6, 7, 8}
	//fmt.Println(a[:])

	//a := []int{6, 7, 8}
	//fmt.Println(a[0:0])

	//var sliceStr []string // 声明字符串切片
	//var sliceNum []int // 声明整型切片
	//var emptySliceNum = []int{} // 声明一个空切片
	//fmt.Println(sliceStr, sliceNum, emptySliceNum) // 输出3个切片
	//fmt.Println(len(sliceStr), len(sliceNum), len(emptySliceNum)) // 输出3个切片大小
	//fmt.Println(sliceStr == nil) // 切片判定空的结果
	//fmt.Println(sliceNum == nil)
	//fmt.Println(emptySliceNum == nil)

	slice1 := make([]int, 6)
	slice2 := make([]int, 6, 10)
	fmt.Println(slice1, slice2)
	fmt.Println(len(slice1), len(slice2))
}
