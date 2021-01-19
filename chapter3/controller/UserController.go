//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package controller

import (
	"fmt"
	"gitee.com/shirdonl/goWebActualCombat/chapter3/model"
	"html/template"
	"net/http"
	"strconv"
)

type UserController struct {
}

func (c UserController) GetUser(w http.ResponseWriter, r *http.Request)  {
	query := r.URL.Query()
	uid, _ := strconv.Atoi(query["uid"][0])

	//此处调用模型从数据库中获取数据
	user := model.GetUser(uid)
	fmt.Println(user)

	t, _ := template.ParseFiles("view/t3.html")
	userInfo := []string{user.Name, user.Phone}
	t.Execute(w, userInfo)
}