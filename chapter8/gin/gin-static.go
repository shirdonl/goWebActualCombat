//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.Static("/assets", "/var/www/gin/assets")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	// 启动服务
	router.Run(":8080")
}
