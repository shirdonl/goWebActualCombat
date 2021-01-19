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
	"strings"
)

func request(w http.ResponseWriter, r *http.Request) {
	//这些信息是输出到服务器端的打印信息
	fmt.Println("Request解析")
	//HTTP方法
	fmt.Println("method", r.Method)
	// RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
	fmt.Println("RequestURI:", r.RequestURI)
	//URL类型,下方分别列出URL的各成员
	fmt.Println("URL_path", r.URL.Path)
	fmt.Println("URL_RawQuery", r.URL.RawQuery)
	fmt.Println("URL_Fragment", r.URL.Fragment)
	//协议版本
	fmt.Println("proto", r.Proto)
	fmt.Println("protomajor", r.ProtoMajor)
	fmt.Println("protominor", r.ProtoMinor)
	//HTTP请求的头域
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Println("header key:" + k + "  value:" + vv)
		}
	}
	//判断是否multipart方式
	isMultipart := false
	for _, v := range r.Header["Content-Type"] {
		if strings.Index(v, "multipart/form-data") != -1 {
			isMultipart = true
		}
	}
	//解析body
	if isMultipart == true {
		r.ParseMultipartForm(128)
		fmt.Println("解析方式:ParseMultipartForm")
	} else {
		r.ParseForm()
		fmt.Println("解析方式:ParseForm")
	}
	//body内容长度
	fmt.Println("ContentLength", r.ContentLength)
	//是否在回复请求后关闭连接
	fmt.Println("Close", r.Close)
	//HOSt
	fmt.Println("host", r.Host)
	//该请求的来源地址
	fmt.Println("RemoteAddr", r.RemoteAddr)

	fmt.Fprintf(w, "hello, let's go!") //这个写入到w的是输出到客户端的
}
func main() {
	http.HandleFunc("/hello", request)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
