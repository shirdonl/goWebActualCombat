//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package pipeline

import (
	"github.com/PuerkitoBio/goquery"
)

type FilePipeLine struct {
	dir string
}

func NewFilePipeLine(dir string) *FilePipeLine{
	return &FilePipeLine{dir:dir}
}

func (p *FilePipeLine) Process(doc *goquery.Document){
	// 文件写入实现
}
