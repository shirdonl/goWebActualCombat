//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package spider

import (
	"downloader"
	"github.com/PuerkitoBio/goquery"
	"log"
	"pageprocess"
	"pipeline"
	"scheduler"
	"strconv"
	"time"
)

// threadnum   - 线程数量
// scheduler   - 调度器
// downloader  - 下载器
// pageprocess - 页面处理
// pipeline    - 输出
type Spider struct {
	threadnum uint8
	scheduler scheduler.Scheduler
	downloader downloader.DownLoader
	pageprocess pageprocess.PageProcess
	pipeline pipeline.PipeLine
}
// NewSpider 创建一个爬虫引擎
func NewSpider(threadnum int,path string) *Spider{
	return &Spider{
		scheduler:scheduler.NewQueueSCheduler(),
		downloader:downloader.NewHttpDownLoader(),
		pageprocess:pageprocess.NewPageProcess(),
		pipeline:pipeline.NewFilePipeLine(path),
		threadnum:uint8(threadnum),
	}
}
// Run 引擎运行
func (s *Spider) Run(){
	// Process并发数量
	rm := NewResourceManagerChan(s.threadnum)
	log.Println("[Spider] 爬虫运行 - 处理线程数:" + strconv.Itoa(rm.Cap()))
	for{
		url,ok := s.scheduler.Pop()
		// 爬取队列为空 并且 没有Process线程在处理 认为爬虫结束
		if ok == false && rm.Has() == 0{
			log.Println("[Spider] 爬虫运行结束")
			break
		}else if ok == false{ // Process线程正在处理，可能还会有新的请求加入调度
			log.Println("[Spider] 爬取队列为空 - 等待处理")
			time.Sleep(500 * time.Millisecond)
			continue
		}
		// 控制Process线程并发数量
		rm.GetOne()
		go func(url string) {
			defer rm.FreeOne()
			s.Process(url)
		}(url)
	}
}
// 添加请求链接
func (s *Spider) AddUrl(url string) *Spider{
	s.scheduler.Push(url)
	return s
}
func (s *Spider) AddUrls(urls []string) *Spider{
	for _,url := range urls{
		s.scheduler.Push(url)
	}
	return s
}
// 处理请求链接
func (s *Spider) Process(url string){
	// 下载链接
	resp := s.downloader.DownLoad(url)
	if resp == nil{
		/*下载失败重新加入调度队列中*/
		if !s.downloader.Visited(url){
			s.scheduler.Push(url)
		}
		return
	}
	// 页面处理 - 使用goquery包简单处理
	doc,err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil{
		log.Println("[Process] 解析错误")
		s.scheduler.Push(url)
		return
	}
	// 将新请求链接加入到调度器中
	links := s.pageprocess.Process(doc)
	for _,url := range links{
		if !s.downloader.Visited(url){
			s.scheduler.Push(url)
		}
	}
	// 输出文档
	go s.pipeline.Process(doc)
}
// 控制线程并发数
