//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

//CREATE TABLE `beego_user` (
// `id` bigint(20) NOT NULL AUTO_INCREMENT,
// `name` varchar(20) DEFAULT '',
// `phone` varchar(20) DEFAULT '',
// PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4

package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)  //数据库类型设计
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/chapter4?charset=utf8")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(BeegoUser))
}

type BeegoUser struct {
	Id   int //默认主健为id
	Name  string
	Phone string
}

func main() {
	o := orm.NewOrm()
	//o.Using("chapter4") // 默认使用 default，你可以指定为其他数据库
	//user := new(BeegoUser)
	//user.Name = "Shirdon2"
	//user.Phone = "18888888888"
	//fmt.Println(o.Insert(user))
	//
	//orm.Debug = true

	//var w io.Writer
	//orm.DebugLog = orm.NewLog(w)



	//查询+++++++
	//user := BeegoUser{}
	//// 先对主键id赋值, 查询数据的条件就是where id=6
	//user.Id = 6
	//
	//// 通过Read函数查询数据
	//// 等价sql: select * from beego_user where id = 6
	//err := o.Read(&user)
	//
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	fmt.Println(user.Id, user.Name)
	//}
	//查询+++++++



	//删除++++++++++
	//user := BeegoUser{}
	//// 先对主键id赋值, 查询数据的条件就是where id=7
	//user.Id = 7
	//
	//if num, err := o.Delete(&user); err != nil {
	//	fmt.Println("删除失败")
	//} else {
	//	fmt.Println("删除数据影响的行数:", num)
	//}
	//删除++++++++++


	//更新++++++++++
	//user := BeegoUser{}
	//// 先对主键id赋值, 查询数据的条件就是where id=7
	//user.Id = 6
	//user.Name = "James"
	//
	//num, err := o.Update(&user)
	//if err != nil {
	//	fmt.Println("更新失败")
	//} else {
	//	fmt.Println("更新数据影响的行数:", num)
	//}
	//更新++++++++++

	//事务++++++++++
	o.Begin()
	user1 := BeegoUser{}
	// 赋值
	user1.Id = 6
	user1.Name = "James"

	user2 := BeegoUser{}
	// 赋值
	user2.Id = 12
	user2.Name = "Wade"

	_, err1 := o.Update(&user1)
	_, err2 := o.Insert(&user2)
	// 检测事务执行状态
	if err1 != nil || err2 != nil {
		// 如果执行失败，回滚事务
		o.Rollback()
	} else {
		// 任务执行成功，提交事务
		o.Commit()
	}
	//事务++++++++++

	//	//查询数据
	//created, id, err := o.ReadOrCreate(&user2, "Id")
	//println(created)
	//println(id)
	//println(err)


	//
	//db, err := orm.GetDB()
	//if err != nil {
	//	fmt.Println("get default database")
	//}
	//
	//db2, err := orm.GetDB("alias")
	//if err != nil {
	//	fmt.Println("get alias database")
	//}

	//o := orm.NewOrm()
	//var r orm.RawSeter
	//r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "jack", "jim")

	type Driver interface {
		Name() string
		Type() int
	}

	//orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/chapter4?charset=utf8")
	//orm.RegisterDataBase("db1", "mysql", "root:root@/orm_db2?charset=utf8")
	//orm.RegisterDataBase("db2", "sqlite3", "data.db")
	//
	//o1 := orm.NewOrm()
	//o1.Using("default")
	//dr := o1.Driver()
	//fmt.Println(dr.Name() == "default") // true
	//fmt.Println(dr.Type() == orm.DRMySQL) // true
	//
	//o2 := orm.NewOrm()
	//o2.Using("db2")
	//dr = o2.Driver()
	//fmt.Println(dr.Name() == "db2") // true
	//fmt.Println(dr.Type() == orm.DRSqlite) // true
}