//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

import "fmt"

func SendNotification(user string) chan string {

	//......此处省略查询数据库获取新消息。
	//声明一个通道来保存消息
	notifications := make(chan string, 500)

	// 开启一个通道
	go func() {
		//将消息放入通道
		notifications <- fmt.Sprintf("Hi %s, welcome to our site!", user)
	}()

	return notifications
}

func main() {
	barry := SendNotification("barry")     //  获取barry的消息
	shirdon := SendNotification("shirdon") // 获取shirdon的消息

	// 获取消息的返回
	fmt.Println(<-barry)
	fmt.Println(<-shirdon)
}
