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
	"reflect"
	"runtime"
	"time"
)

func main() {
	http.HandleFunc("/", log(index))
	http.ListenAndServe(":8087", nil)
}

func log2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//...
		h.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello index!!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(" time: %s| handlerfunc:%s\n", time.Now().String(), runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name())
		h(w, r)
	}
}
