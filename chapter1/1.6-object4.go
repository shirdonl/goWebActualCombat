//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

type Engine interface {
	Start()
	Stop()
}

type Bus struct {
	Engine // 包含 Engine 类型的匿名字段
}

func (c *Bus) Working() {
	c.Start() //开动汽车
	c.Stop()  //停车
}
