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
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

var (
	db *gorm.DB
	sqlConnection = "root:a123456@tcp(127.0.0.1:3306)/chapter8?" +
		"charset=utf8&parseTime=true"
)

//初始化
func init() {
	//打开数据库连接
	var err error
	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
}

func main() {
	router := gin.Default()
	v2 := router.Group("/api/v2/user")
	{
		v2.POST("/", createUser)      //POST方法，创建新用户
		v2.GET("/", fetchAllUser)     //GET方法，获取所有用户
		v2.GET("/:id", fetchUser)     //GET方法，获取某一个用户，形如：/api/v2/user/1
		v2.PUT("/:id", updateUser)    //PUT方法，更新用户，形如：/api/v2/user/1
		v2.DELETE("/:id", deleteUser) //DELETE方法，删除用户，形如：/api/v2/user/1
	}
	router.Run("127.0.0.1:8086")
}

type (
	// 数据表结构体类
	User struct {
		ID       uint   `json:"id"`
		Phone    string `json:"phone"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	//响应返回的结构体
	UserRes struct {
		ID    uint   `json:"id"`
		Phone string `json:"phone"`
		Name  string `json:"name"`
	}
)

//md5加密
func md5Password(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 创建新用户
func createUser(c *gin.Context) {
	phone := c.PostForm("phone") //获取POST请求参数phone
	name := c.PostForm("name")   //获取POST请求参数name
	user := User{
		Phone:    phone,
		Name:     name,
		Password: md5Password("666666"), //用户密码
	}
	db.Save(&user) //保存到数据库
	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "User created successfully!",
			"ID":      user.ID,
		})         //返回状态到客户端
}

// 获取所有用户
func fetchAllUser(c *gin.Context) {
	var user []User        //定义一个数组去数据库总获取数据
	var _userRes []UserRes //定义一个响应数组用户返回数据到客户端

	db.Find(&user)

	if len(user) <= 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No user found!",
			})
		return
	}

	//循环遍历，追加到响应数组
	for _, item := range user {
		_userRes = append(_userRes,
			UserRes{
				ID:    item.ID,
				Phone: item.Phone,
				Name:  item.Name,
			})
	}
	c.JSON(http.StatusOK,
		gin.H{"status":
		http.StatusOK,
			"data": _userRes,
		}) //返回状态到客户端
}

// 获取单个用户
func fetchUser(c *gin.Context) {
	var user User       //定义User结构体
	ID := c.Param("id") //获取参数id

	db.First(&user, ID)

	if user.ID == 0 { //如果用户不存在，则返回
		c.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	//返回响应结构体
	res := UserRes{ID: user.ID, Phone: user.Phone, Name: user.Name}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": res})
}

//更新用户
func updateUser(c *gin.Context) {
	var user User           //定义User结构体
	userID := c.Param("id") //获取参数id
	db.First(&user, userID) //查找数据库

	if user.ID == 0 { //如果数据库不存在，则返回
		c.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	//更新对应的字段值
	db.Model(&user).Update("phone", c.PostForm("phone"))
	db.Model(&user).Update("name", c.PostForm("name"))
	c.JSON(http.StatusOK,
		gin.H{"status": http.StatusOK, "message": "Updated User successfully!"})
}

// 删除用户
func deleteUser(c *gin.Context) {
	var user User           //定义User结构体
	userID := c.Param("id") //获取参数id

	db.First(&user, userID) //查找数据库

	if user.ID == 0 { //如果数据库不存在，则返回
		c.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	//删除用户
	db.Delete(&user)
	c.JSON(http.StatusOK,
		gin.H{"status": http.StatusOK, "message": "User deleted successfully!"})
}
