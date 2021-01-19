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

type Books struct {
	title string
	author string
	subject string
	press string
}

func main() {
	var bookGo Books        //声明bookGo为Books类型
	var bookPython Books        //声明bookPython为Books类型

	// bookGo描述
	bookGo.title = "Go Web编程实战派从入门到精通"
	bookGo.author = "廖显东"
	bookGo.subject = "Go语言教程"
	bookGo.press = "电子工业出版社"

	// bookPython描述
	bookPython.title = "Python教程xxx"
	bookPython.author = "张三"
	bookPython.subject = "Python语言教程"
	bookPython.press = "xxx出版社"

	//打印 bookGo 信息
	fmt.Printf( "bookGo title : %s\n", bookGo.title)
	fmt.Printf( "bookGo author : %s\n", bookGo.author)
	fmt.Printf( "bookGo subject : %s\n", bookGo.subject)
	fmt.Printf( "bookGo press : %s\n", bookGo.press)

	//打印 bookPython 信息
	fmt.Printf( "bookPython title : %s\n", bookPython.title)
	fmt.Printf( "bookPython author : %s\n", bookPython.author)
	fmt.Printf( "bookPython subject : %s\n", bookPython.subject)
	fmt.Printf( "bookPython press : %s\n", bookPython.press)
}