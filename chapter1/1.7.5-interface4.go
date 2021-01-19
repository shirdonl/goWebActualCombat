package main

import "fmt"

// message是一个定义了通知类行为的接口
type Message interface {
	sending()
}

// 定义user及user.notify方法
type User struct {
	name  string
	phone string
}

func (u *User) sending() {
	fmt.Printf("Sending user phone to %s<%s>\n", u.name, u.phone)
}

// 定义admin及admin.message方法
type admin struct {
	name  string
	phone string
}

func (a *admin) sending() {
	fmt.Printf("Sending admin phone to %s<%s>\n", a.name, a.phone)
}

func main() {
	// 创建一个user值并传给sendMessage
	bill := User{"Barry", "barry@gmail.com"}
	sendMessage(&bill)

	// 创建一个admin值并传给sendMessage
	lisa := admin{"Jim", "jim@gmail.com"}
	sendMessage(&lisa)
}

// sendMessage接受一个实现了message接口的值 并发送通知
func sendMessage(n Message) {
	n.sending()
}
