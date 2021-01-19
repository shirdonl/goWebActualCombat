//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package spider

type ResourceManager struct {
	tc chan uint8
}

func NewResourceManagerChan(num uint8) *ResourceManager{
	tc := make(chan uint8,num)
	return &ResourceManager{tc:tc}
}

func (r *ResourceManager) GetOne(){
	r.tc <- 1
}

func (r *ResourceManager) FreeOne(){
	<- r.tc
}

func (r *ResourceManager) Cap() int{
	return cap(r.tc)
}

func (r *ResourceManager) Has() int{
	return len(r.tc)
}

func (r *ResourceManager) Left() int{
	return cap(r.tc) - len(r.tc)
}
