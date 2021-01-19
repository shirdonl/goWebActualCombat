package main

import "fmt"

func main() {
	name := "go"
	defer fmt.Println(name) // 输出: go

	name = "python"
	defer fmt.Println(name) // 输出: python

	name = "java"
	fmt.Println(name)
}