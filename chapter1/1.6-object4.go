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
