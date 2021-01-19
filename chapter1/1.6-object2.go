//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

//package main
//
//// 三角形结构体
//type Student struct {
//	Name  string
//	score float32
//}
//
//func main() {
//	student := new(Student)
//	student.Name = "shirdon"
//	student.score = 9
//
//	println(student.Name)
//	println(student.score)
//}

package main

import (
	"fmt"
	"gitee.com/shirdonl/goWebActualCombat/chapter1/person"
)

func main() {
	s := new(person.Student)
	s.SetName("Shirdon")
	fmt.Println(s.GetName())
}
