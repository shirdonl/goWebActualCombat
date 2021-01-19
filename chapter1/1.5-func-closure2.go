package main

import "fmt"

// 遍历切片中每个元素，通过给定的函数访问元素
func visitPrint(list []int, f func(int)) {
	for _, value := range list {
		f(value)
	}
}

func main() {
	sli := []int{1, 6, 8}
	// 使用匿名函数打印切片的内容
	visitPrint(sli, func(value int) {
		fmt.Println(value)
	})
}
