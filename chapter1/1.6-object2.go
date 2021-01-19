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
