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
	"io"
	"net/http"
	"os"
	"strconv"
)

func Get(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	//读取网页的body内容
	buf := make([]byte, 4*1024)
	for true {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			} else {
				fmt.Println("resp.Body.Read err = ", err)
				break
			}
		}
		result += string(buf[:n])
	}
	return
}

//将所有的网页内容爬取下来
func SpiderPage(i int, page chan<- int) {
	url := "https://github.com/search?q=go&type=Repositories&p=1" + strconv.Itoa((i-1)*50)
	fmt.Printf("正在爬取第%d个网页\n", i)
	//爬,将所有的网页内容爬取下来
	result, err := Get(url)
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}
	//把内容写入到文件
	filename := "page"+strconv.Itoa(i) + ".html"
	f, err1 := os.Create(filename)
	if err1 != nil {
		fmt.Println("os.Create err = ", err1)
		return
	}
	//写内容
	f.WriteString(result)
	//关闭文件
	f.Close()
	//每爬完一个，就给个值
	page <- i
}

func Run(start, end int) {
	fmt.Printf("正在爬取第%d页到%d页\n", start, end)
	//因为很有可能爬虫还没有结束下面的循环就已经结束了，所以这里就需要且到通道
	page := make(chan int)
	for i := start; i <= end; i++ {
		//将page阻塞
		go SpiderPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page) //这里直接将面码传给点位符，值直接从管道里取出
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页数字>=1：> ")
	fmt.Scan(&start)
	fmt.Printf("请输入结束页数字：> ")
	fmt.Scan(&end)
	Run(start, end)
}
