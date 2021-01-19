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
	"time"
)

func cookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "cookie_id", Value: "ALNI_MZXW8Ls0vGogdVKX1BW4JA0Tzszgw", Expires: expiration}
	http.SetCookie(w, &cookie)
	//fmt.Fprintln(w, cookie)
	c, _ := r.Cookie("cookie_id")
	fmt.Fprint(w, c)
}

func main() {
	http.HandleFunc("/", cookie)  //使用http包默认的多路复用器实例绑定
	http.ListenAndServe("127.0.0.1:8087", nil)  //使用http包默认的服务器实例，并开启监听
}
