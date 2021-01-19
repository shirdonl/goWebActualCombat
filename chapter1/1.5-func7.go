package main

import "fmt"

func main() {
	x, y := 6, 8
	defer func(a int) {
		fmt.Println("defer x, y = ", a, y) //y为闭包引用
	}(x)
	x += 10
	y += 100
	fmt.Println(x, y)
}
