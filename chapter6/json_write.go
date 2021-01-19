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
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	UserName string
	NickName string `json:"nickname"`
	Email    string
}

func main() {
	user := &User{
		UserName: "Jack",
		NickName: "Ma",
		Email:    "xxxxx@qq.com",
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("json.Marshal failed,err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))

	file, _ := os.Create("json_write.json")
	defer file.Close()
	file.Write(data)
}
