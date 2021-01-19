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
	"math"
)

func main() {
	getPow := func(x float64, y float64) float64 { //声明函数变量
		return math.Pow(x, y)
	}
	fmt.Println(getPow(3, 4)) //使用函数
}
