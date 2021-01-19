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
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
)

var db *sql.DB

type User struct {
	Uid      int
	Name     string
	Phone    string
	Email    string
	Password string
}

//定义一个全局变量
var u User

//初始化数据库连接
func init() {
	db, _ = sql.Open("mysql",
		"root:a123456@tcp(127.0.0.1:3306)/chapter6?"+
			"charset=utf8mb4&parseTime=True&loc=Local")
}

func main() {
	//定义导出的文件名
	filename := "./exportUsers.csv"

	//从数据库中获取数据
	users := queryMultiRow()
	//定义一个二维数组
	column := [][]string{{"手机号", "用户UID", "Email", "用户名"}}
	for _, u := range users {
		str := []string{}
		str = append(str, u.Phone)
		str = append(str, strconv.Itoa(u.Uid))
		str = append(str, u.Email)
		str = append(str, u.Name)
		column = append(column, str)
	}
	//导出
	ExportCsv(filename, column)
}

//导出csv文件
func ExportCsv(filePath string, data [][]string) {
	fp, err := os.Create(filePath) //创建文件句柄
	if err != nil {
		log.Fatalf("创建文件["+filePath+"]句柄失败,%v", err)
		return
	}
	defer fp.Close()
	fp.WriteString("\xEF\xBB\xBF") //写入UTF-8 BOM
	w := csv.NewWriter(fp)         //创建一个新的写入文件流
	w.WriteAll(data)
	w.Flush()
}

// 查询多条数据
func queryMultiRow() ([]User) {
	rows, err := db.Query("select uid,name,phone,email from `user` where uid > ?", 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	users := []User{}
	for rows.Next() {
		err := rows.Scan(&u.Uid, &u.Name, &u.Phone, &u.Email)
		users = append(users, u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
	}
	return users
}
