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
	"reflect"
)

func main() {
	var name string = "Go Web Program"
	v1 := reflect.ValueOf(&name)
	fmt.Println("v1 可写性为:", v1.CanSet())
	v2 := v1.Elem()
	fmt.Println("v2 可写性为:", v2.CanSet())
}
