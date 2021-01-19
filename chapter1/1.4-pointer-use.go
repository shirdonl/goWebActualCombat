package main

import (
	"fmt"
)

func main() {
	var score int = 100
	var name string = "Barry"
	// 使用 fmt.Printf 的动词%p输出 score 和 name 变量取地址后的指针值
	fmt.Printf("%p %p", &score, &name)
}
