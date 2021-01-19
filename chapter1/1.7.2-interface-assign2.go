package main

import (
	"gitee.com/shirdonl/goWebActualCombat/chapter1/oop1"
	"gitee.com/shirdonl/goWebActualCombat/chapter1/oop2"
)

type Num int

func (x Num) Equal(i int) bool {
	return int(x) == i
}

func (x Num) LessThan(i int) bool {
	return int(x) < i
}

func (x Num) BiggerThan(i int) bool {
	return int(x) > i
}

func (x *Num) Sum(i int) {
	*x = *x + Num(i)
}

func main() {
	//var f1 Num = 6
	//var f2 oop1.NumInterface1 = f1
	//var f3 oop2.NumInterface2 = f2
	//println(f3)

	//var f1 Num = 6
	//var f2 oop2.NumInterface2 = f1
	//var f3 oop1.NumInterface1 = f2

}
