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

type Num int

func (x Num) Equal(i Num) bool {
	return x == i
}

func (x Num) LessThan(i Num) bool {
	return x < i
}

func (x Num) BiggerThan(i Num) bool {
	return x > i
}

func (x *Num) Multiple(i Num) {
	*x = *x * i
}

func (x *Num) Divide(i Num) {
	*x = *x / i
}

type NumI interface {
	Equal(i Num) bool
	LessThan(i Num) bool
	BiggerThan(i Num) bool
	Multiple(i Num)
	Divide(i Num)
}

type NumI2 interface {
	Equal(i Num) bool
	LessThan(i Num) bool
	MoreThan(i Num) bool
}

func main() {
	var x Num = 8
	var y NumI = &x

	var a Num = 1
	var b1 NumI = &a
	var b2 NumI2 = a
	fmt.Println(b1)
	fmt.Println(b2)
}
