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
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	html := `<html> 
        <head>
            <title>Write方法返回HTML文档</title>
        </head> 
        <body>
            <h1>你好，欢迎一起学习《Go Web编程实战派从入门到精通》
        </body> 
    </html>`
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", Home)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Println(err)
	}
}
