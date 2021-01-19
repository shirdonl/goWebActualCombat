//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

type stringer interface {
	string() string
}

type tester interface {
	stringer // 嵌入其他接口
	test()
}

type data struct{}

func (*data) test() {
}

func (data) string() string {
	return ""
}

func main() {
	var d data

	var t tester = &d
	t.test()
	println(t.string())
}
