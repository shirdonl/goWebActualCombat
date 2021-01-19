package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	press   string
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Book{"Go Web编程实战派从入门到精通", "廖显东", "Go语言教程", "电子工业出版社"})
	// 也可以使用 key => value 格式
	fmt.Println(Book{title: "Go Web编程实战派从入门到精通", author: "廖显东", subject: "Go语言教程", press: "电子工业出版社"})
	// 忽略的字段为 0 或 空
	fmt.Println(Book{title: "Go Web编程实战派从入门到精通", author: "廖显东"})
}
