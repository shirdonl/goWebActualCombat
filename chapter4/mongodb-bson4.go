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
	"context"
	"gitee.com/shirdonl/goWebActualCombat/chapter4/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type DeleteCond struct {
	BeforeCond TimeBeforeCond `bson:"tp.startTime"`
}

//startTime小于某时间，使用这种方式可以对想要进行的操作($set、$group等)提前定义
type TimeBeforeCond struct {
	BeforeTime int64 `bson:"$lt"`
}

func main() {
	var (
		client     = mongodb.MgoCli()
		collection *mongo.Collection
		err        error
		uResult    *mongo.DeleteResult
		delCond    *DeleteCond
	)
	collection = client.Database("my_db").Collection("table1")

	//删除jobName为job0的数据
	delCond = &DeleteCond{
		BeforeCond: TimeBeforeCond{
			BeforeTime: time.Now().Unix()}}
	if uResult, err = collection.DeleteMany(context.TODO(),
		delCond); err != nil {
		log.Fatal(err)
	}
	log.Println(uResult.DeletedCount)
}