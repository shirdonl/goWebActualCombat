//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package pageprocess

import (
	"github.com/PuerkitoBio/goquery"
)

type PageProcess struct {
}

func NewPageProcess() PageProcess{
	return PageProcess{}
}

// 返回链接函数
func (p *PageProcess) Process(d *goquery.Document) []string{
	var links []string
	// 获取链接的处理代码
	return links
}
