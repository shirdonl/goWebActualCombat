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
	"html/template"
	"net/http"
)

//定义一个UserInfo结构体
type UserInfo struct {
	Name string
	Gender string
	Age int
}

func tmplSample(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.html", "./ul.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "张三",
		Gender: "男",
		Age:    28,
	}
	tmpl.Execute(w, user)
	fmt.Println(tmpl)
}

func main() {
	http.HandleFunc("/", tmplSample)
	http.ListenAndServe(":8087", nil)
}
