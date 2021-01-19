package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	press   string
}

func main() {
	var bookGo Books     /* 声明bookGo为Books类型 */
	var bookPython Books /* 声明bookPython为Books类型 */

	/* bookGo描述 */
	bookGo.title = "Go Web编程实战派从入门到精通"
	bookGo.author = "廖显东"
	bookGo.subject = "Go语言教程"
	bookGo.press = "电子工业出版社"

	/* bookPython描述 */
	bookPython.title = "Python教程xxx"
	bookPython.author = "张三"
	bookPython.subject = "Python语言教程"
	bookPython.press = "xxx出版社"

	/* 打印 bookPython 信息 */
	printBook(&bookGo)

	/* 打印 bookPython 信息 */
	printBook(&bookPython)
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book press : %s\n", book.press)
}
