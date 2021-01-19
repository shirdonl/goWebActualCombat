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
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	testM := bson.M{
		"jobName": "job multi1",
	}
	var raw bson.Raw
	tmp, _ := bson.Marshal(testM)
	bson.Unmarshal(tmp, &raw)

	fmt.Println(testM)
	fmt.Println(raw)
}