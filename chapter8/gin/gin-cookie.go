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
	router := gin.Default()
	router.GET("/cookie", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("my_cookie", "cookievalue", 3600, "/", "localhost", false, true)
	})

	r.Run(":8068")
}

func Handler(c *gin.Context) {
	// 根据cookie名字读取cookie值
	data, err := c.Cookie("my_cookie")
	if err != nil {
		// 直接返回cookie值
		c.String(200,data)
		return
	}
	c.String(200,"not found!")
}