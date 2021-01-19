//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package main

import "net/url"

func main() {
	path := "http://lcoalhost:8082/article?id=1"
	p, _ := url.Parse(path) // 解析url
	println(p.Host)
	println(p.User)
	println(p.RawQuery)
	println(p.RequestURI())
}
