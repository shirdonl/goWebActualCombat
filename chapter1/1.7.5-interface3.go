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

func main() {
	var a interface{} = func(a int) string {
		return fmt.Sprintf("d:%d", a)
	}
	switch b := a.(type) { // 局部变量b是类型转换后的结果
	case nil:
		println("nil")
	case *int:
		println(*b)
	case func(int) string:
		println(b(66))
	case fmt.Stringer:
		fmt.Println(b)
	default:
		println("unknown")
	}
}
