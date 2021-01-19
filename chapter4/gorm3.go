package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var (
	db            *gorm.DB
	sqlConnection = "root:a123456@tcp(127.0.0.1:3306)/chapter4?" +
		"charset=utf8&parseTime=true"
)

// 数据表结构体类
type GormUser struct {
	ID       uint   `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

//初始化
func init() {
	//打开数据库连接
	var err error
	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&GormUser{})
}

func main() {
	defer db.Close()

	//创建用户
	//GormUser := GormUser{
	//	Phone:    "13888888888",
	//	Name:     "Shirdon",
	//	Password: md5Password("666666"), //用户密码
	//}
	//db.Save(&GormUser) //保存到数据库
	//db.Create(&GormUser) //保存到数据库

	//查询用户
	//var GormUser = new(GormUser)
	//db.Where("phone = ?", "18888888888").Find(&GormUser)
	//db.First(&GormUser, "phone = ?", "18888888888")
	//fmt.Println(GormUser)
	//
	////更新用户
	//var GormUser = new(GormUser)
	//err:=db.Model(&GormUser).Where("phone = ?", "18888888888").
	//	Update("phone", "13888888888").Error
	//if err !=nil {
	//	//
	//}
	//
	////删除用户
	//var GormUser = new(GormUser)
	//db.Where("phone = ?", "13888888888").Delete(&GormUser)


	////开启事务
	//tx := db.Begin()
	//
	//GormUser := GormUser{
	//	Phone:    "18888888888",
	//	Name:     "Shirdon",
	//	Password: md5Password("666666"), //用户密码
	//}
	//if err := tx.Create(&GormUser).Error; err != nil {
	//	//事务回滚
	//	tx.Rollback()
	//	fmt.Println(err)
	//}
	//db.First(&GormUser, "phone = ?", "18888888888")
	////事务提交
	//tx.Commit()


	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))


}

//md5加密
func md5Password(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
