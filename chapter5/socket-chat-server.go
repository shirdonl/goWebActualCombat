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
	"net"
	"time"
)

var ConnSlice map[net.Conn]*Heartbeat

type Heartbeat struct {
	endTime int64 //过期时间
}

func main() {
	ConnSlice = map[net.Conn]*Heartbeat{}
	l, err := net.Listen("tcp", "127.0.0.1:8086")
	if err != nil {
		fmt.Println("服务器启动失败")
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
		}
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		ConnSlice[conn] = &Heartbeat{
			endTime: time.Now().Add(time.Second * 5).Unix(), //初始化过期时间
		}
		go handelConn(conn)
	}
}
func handelConn(c net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := c.Read(buffer)
		if ConnSlice[c].endTime > time.Now().Unix() {
			ConnSlice[c].endTime = time.Now().Add(time.Second * 5).Unix() //更新心跳时间
		} else {
			fmt.Println("长时间未发消息断开连接")
			return
		}
		if err != nil {
			return
		}
		//如果是心跳检测，那就不要执行剩下的代码
		if string(buffer[0:n]) == "1" {
			c.Write([]byte("1"))
			continue
		}
		for conn, heart := range ConnSlice {
			if conn == c {
				continue
			}
			//心跳检测 使用懒惰更新,需要发送数据的时候才检查规定时间内有没有数据到达
			if heart.endTime < time.Now().Unix() {
				delete(ConnSlice, conn) //从房间列表中删除连接，并且关闭
				conn.Close()
				fmt.Println("删除连接", conn.RemoteAddr())
				fmt.Println("现在存有链接", ConnSlice)
				continue
			}
			conn.Write(buffer[0:n])
		}
	}
}
