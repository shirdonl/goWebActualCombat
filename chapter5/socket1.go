//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package main

import "net"

func main()  {
	conn, err := net.Dial("tcp", "192.168.0.1:8087")
	conn, err := net.Dial("udp", "192.168.0.2:8088")
	conn, err := net.Dial("ip4:icmp", "www.shirdon.com")
	net.DialTCP()
	println(conn)
	println(err)
}
