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
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]//从参数中读取主机信息
	conn, err := net.Dial("tcp", service)//建立网络连接
	validateError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))//调用返回的连接对象提供的 Write 方法发送请求
	validateError(err)

	result, err := fullyRead(conn)
	validateError(err)

	fmt.Println(string(result))// 打印响应数据

	os.Exit(0)
}

//连接出错则打印错误消息并退出程序
func validateError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//通过连接对象提供的 Read 方法读取所有响应数据
func fullyRead(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}