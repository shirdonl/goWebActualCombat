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
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type EmailConfig struct {
	XMLName  xml.Name `xml:"config"`
	SmtpServer string `xml:"smtpServer"`
	SmtpPort int `xml:"smtpPort"`
	Sender string `xml:"sender"`
	SenderPassword string `xml:"senderPassword"`
	Receivers EmailReceivers `xml:"receivers"`
}

type EmailReceivers struct {
	Flag string `xml:"flag,attr"`
	User []string `xml:"user"`
}

func main() {
	file, err := os.Open("email_config.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := EmailConfig{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
	fmt.Println("SmtpServer is : ",v.SmtpServer)
	fmt.Println("SmtpPort is : ",v.SmtpPort)
	fmt.Println("Sender is : ",v.Sender)
	fmt.Println("SenderPasswd is : ",v.SenderPassword)
	fmt.Println("Receivers.Flag is : ",v.Receivers.Flag)
	for i,element := range v.Receivers.User {
		fmt.Println(i,element)
	}
}