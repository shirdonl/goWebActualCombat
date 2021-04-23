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
	"io/ioutil"
	"net/http"
)

func Welcome() string { //没参数
	return "Welcome"
}

func Doing(name string) string { //有参数
	return name + ", Learning Go Web template "
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("./funcs.html")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}
	// 自定义一个匿名模板函数
	loveGo := func() (string) {
		return "欢迎一起学习《Go Web编程实战派从入门到精通》"
	}
	// 采用链式操作在Parse()方法之前调用Funcs添加自定义的loveGo函数
	tmpl1, err := template.New("funcs").Funcs(template.FuncMap{"loveGo": loveGo}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	funcMap := template.FuncMap{
		//在FuncMap中声明相应要使用的函数，然后就能够在template字符串中使用该函数
		"Welcome": Welcome,
		"Doing":   Doing,
	}
	name := "Shirdon"
	tmpl2, err := template.New("test").Funcs(funcMap).Parse("{{Welcome}}\n{{Doing .}}\n")
	if err != nil {
		panic(err)
	}

	// 使用user渲染模板，并将结果写入w
	tmpl1.Execute(w, name)
	tmpl2.Execute(w, name)
}

func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":8087", nil)

}
