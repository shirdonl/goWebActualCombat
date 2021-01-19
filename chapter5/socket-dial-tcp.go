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
	"io/ioutil"
	"net"
	"os"
)

func main() {
	service := "127.0.0.1:8086"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	fmt.Println("tcpAddr :")
	typeof(tcpAddr)

	myConn, err1 := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err1)
	fmt.Println("myConn :")
	typeof(myConn)

	_, err = myConn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(myConn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func typeof(v interface{}) {
	fmt.Printf("type is:%T\n", v)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
