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

func main() {
	var address = "Chengdu,China"         // 准备一个字符串类型
	ptr := &address                       // 对字符串取地址, ptr类型为*string
	fmt.Printf("ptr type: %T\n", ptr)     // 打印ptr的类型
	fmt.Printf("address: %p\n", ptr)      // 打印ptr的指针地址
	value := *ptr                         // 对指针进行取值操作
	fmt.Printf("value type: %T\n", value) // 取值后的类型
	fmt.Printf("value: %s\n", value)      // 指针取值后就是指向变量的值
}
