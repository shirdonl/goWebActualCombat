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
	"log"
	"net/http"
	"regexp"
)
func index(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	}
	fmt.Fprintf(w, "This words are written by Go") //这个写入到w的是输出到客户端的
}
func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		fmt.Println("method:", r.Method) //获取请求的方法
		//数字匹配,如果不是手机号格式，提示手机号错误
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.PostFormValue("phone")); !m {
			//提示手机号错误
			fmt.Fprintf(w, "the phone input is not right")
			return
		}
		//请求的是登陆数据，那么执行登陆的逻辑判断
		if r.PostFormValue("phone") == "13855667788" && r.PostFormValue("password") == "admin" {
			//如果账号密码正确，跳转到首页
			http.Redirect(w, r, "/", 302)
		} else {
			//提示账号密码错误
			fmt.Fprintf(w, "phone or password is not right")
		}
		r.ParseForm()
		fmt.Println("phone:", r.Form["phone"])
		fmt.Println("password:", r.Form["password"])
	}
}
func main() {
	http.HandleFunc("/", index) //设置首页的路由
	http.HandleFunc("/login", login) //设置登陆页的路由
	err := http.ListenAndServe(":8088", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
