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
	"io"
	"log"
	"net/http"
)

func main() {
	// 设置路由
	http.HandleFunc("/v1", version1)
	// 路由注册完，开始运行
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//version1 handler
func version1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is a version 1 response!")
}
