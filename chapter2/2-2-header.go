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
	"net/http"
)

func main() {
	headers := http.Header{"token": {"kuysdfaeg6634fwr324brfh3urhf839hf349h"}}
	headers.Add("Accept-Charset","UTF-8")
	headers.Set("Host","www.shirdon.com")
	headers.Set("Location","www.baidu.com")
}
