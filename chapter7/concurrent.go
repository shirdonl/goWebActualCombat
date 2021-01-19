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
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	DataRoot     = "./tmp/" // 存放封面图的根目录
	TimeoutLimit = 10       // 设置超时时间
)

// 表示章节ID和封面URL的对应关系
type VolumeCover struct {
	VolumeID int
	Url      string
	Lock     sync.Mutex
	Msg      chan string
	Flag     bool "timeout_flag"
}

// 将图片下载并保存到本地
func SaveImage(vc *VolumeCover) {
	res, err := http.Get(vc.Url)
	defer res.Body.Close()
	if err != nil {
		vc.Msg <- (strconv.Itoa(vc.VolumeID) + " HTTP_ERROR")
	}
	// 创建文件
	dst, err := os.Create(DataRoot + strconv.Itoa(vc.VolumeID) + ".jpg")
	if err != nil {
		vc.Msg <- (strconv.Itoa(vc.VolumeID) + " OS_ERROR")
	}
	// 生成文件
	io.Copy(dst, res.Body)
	// goroutine通信
	vc.Lock.Lock()
	vc.Msg <- "in"
	vc.Flag = true
	vc.Lock.Unlock()
}

func Start(name string, password string, limit int) error {
	runtime.GOMAXPROCS(4)
	sl, err := sql.Open("mysql", name+":"+password+"@/xxx?charset=utf8")
	defer sl.Close()
	if err != nil {
		return err
	}
	// 构造SELECT语句并检索
	queryStr := "SELECT VolumeID, ImageUrl FROM volume "
	if limit > 0 {
		queryStr += "limit " + strconv.Itoa(limit)
	}
	rows, err := sl.Query(queryStr)
	defer rows.Close()
	if err != nil {
		return err
	}

	// 构建列表
	i := 0
	vclist := make([]*VolumeCover, limit)
	for rows.Next() {
		vc := &VolumeCover{}
		if err := rows.Scan(&vc.VolumeID, &vc.Url); err != nil {
			return err
		}
		vc.Msg = make(chan string, 1)
		//vc.To = make(chan bool, 1)
		vc.Lock = *new(sync.Mutex)
		vclist[i] = vc
		i++
	}

	// start goroutines
	for i := 0; i < len(vclist); i++ {
		go SaveImage(vclist[i])
		go func(i int) {
			time.Sleep(TimeoutLimit * time.Second)
			if vclist[i].Flag == false {
				vclist[i].Msg <- "out"
			}
		}(i)
	}

	// 阻塞地获取结果
	for i := 0; i < len(vclist); i++ {
		func(c *VolumeCover) {
			select {
			case m := <-c.Msg:
				fmt.Println(m)
			}

		}(vclist[i])
	}
	return nil
}
