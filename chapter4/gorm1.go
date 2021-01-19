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
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User2 用户信息
type User2 struct {
	//gorm.Model
	Id   int
	Name  string
	Phone string
}


func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/chapter4?" +
		"charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 自动迁移
	//db.AutoMigrate(&UserInfo{})

	u1 := &User2{4, "laoshu", "18888888888",}
	//u2 := UserInfo{7,"hunya", "18777777777"}
	// 创建记录
	db.Create(u1)
	//db.Create(&u2)
	// 查询
	//var u = new(UserInfo)
	//db.First(u)
	//fmt.Printf("%#v\n", u)
	//var uu UserInfo
	//db.Find(&uu, "name=?", "laoshu")
	//fmt.Printf("%#v\n", uu)
	//// 更新
	//db.Model(&u).Update("name", "milaoshu")
	//// 删除
	//db.Delete(&u)
}
