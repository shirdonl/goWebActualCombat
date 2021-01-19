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
	"log"
	"net/http"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, Go HandleFunc")
}

type welcomeHandler struct {
	Name string
}

func (h welcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi, %s", h.Name)
}

func main() {
	mux := http.NewServeMux()
	// 注册处理器函数
	mux.HandleFunc("/hi", hiHandler)

	// 注册处理器
	mux.Handle("/welcome/goweb", welcomeHandler{Name: "Hi, Go Handle"})

	server := &http.Server {
		Addr:       ":8085",
		Handler:    mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
