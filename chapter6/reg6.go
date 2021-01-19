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
	"regexp"
)

func main() {
	res2 := findPhoneNumber("13688888888")
	fmt.Println(res2) // true

	res2 = findPhoneNumber("02888888888")
	fmt.Println(res2) // false

	res2 = findPhoneNumber("123456789")
	fmt.Println(res2) // false
}

func findPhoneNumber(str string) bool {
	// 创建一个正则表达式匹配规则对象
	reg := regexp.MustCompile("^1[1-9]{10}")
	// 利用正则表达式匹配规则对象匹配指定字符串
	res := reg.FindAllString(str, -1)
	if (res == nil) {
		return false
	}
	return true
}
