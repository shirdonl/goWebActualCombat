package main

import "fmt"

/* 定义和的函数 */
func sum(x, y int) int {
	return x + y
}

func main() {
	a, b := compute(6, 8)
	fmt.Println(a, b)
}
