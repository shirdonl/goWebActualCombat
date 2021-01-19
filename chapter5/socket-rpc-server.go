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
	"net/rpc"
)

func main() {
	//实例化
	algorithm := new(Algorithm)
	fmt.Println("Algorithm start", algorithm)
	//注册服务
	rpc.Register(algorithm)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8808", nil)
	if err != nil {
		fmt.Println("err=====", err.Error())
	}
}

type Algorithm int

//参数结构体
type Args struct {
	X, Y int
}

//参数结构体
type Response int

//定义一个方法求两个数的和
//该方法的第一个参数为输入参数，第二个参数为返回值
func (t *Algorithm) Sum(args *Args, reply *int) error {
	*reply = args.X + args.Y
	fmt.Println("Exec Sum ", reply)
	return nil
}

