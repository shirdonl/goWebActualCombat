//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

import (
	"fmt"
)

/* 定义结构体 */
type Circle struct {
	radius float64
}

func main() {
	var c Circle
	c.radius = 6.00
	fmt.Println("圆的周长 = ", c.getCircumference())
}

func (c Circle) getCircumference() float64 { //该 method 属于 Circle 类型对象中的方法
	return 2 * 3.14 * c.radius //c.radius 即为 Circle 类型对象中的属性
}
