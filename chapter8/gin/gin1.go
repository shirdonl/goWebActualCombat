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
	r := gin.Default()
	r.GET("/hi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "gin",
		})
		c.Header("site","shirdon")
	})
	r.Run(":8068")
}
