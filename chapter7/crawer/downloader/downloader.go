//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package downloader

import (
	"crypto/md5"
	"log"
	"net/http"
	"sync"
)

type HttpDownLoader struct {
	locker *sync.Mutex
	downloaded map[[md5.Size]byte] bool
}

func NewHttpDownLoader() *HttpDownLoader{
	locker := new(sync.Mutex)
	downloaded := make(map[[md5.Size]byte]bool)
	return &HttpDownLoader{
		locker:locker,
		downloaded:downloaded,
	}
}

// 下载链接
func (h *HttpDownLoader) DownLoad(url string) *http.Response{
	key := md5.Sum([]byte(url))
	resp,err := http.Get(url)
	h.locker.Lock()
	// 已经被访问过了，不需要访问。
	if ok,has := h.downloaded[key]; has && ok{
		h.locker.Unlock()
		return nil
	}
	// 访问失败
	if err != nil || resp.StatusCode != http.StatusOK{
		log.Println("[DownLoader] 下载链接失败:" + url)
		h.downloaded[key] = false
		h.locker.Unlock()
		return nil
	}
	h.downloaded[key] = true
	h.locker.Unlock()
	log.Println("[DownLoader] 下载链接成功:" + url)
	return resp
}

// 链接是否被访问
func (h *HttpDownLoader) Visited(url string) bool{
	key := md5.Sum([]byte(url))
	var ret bool
	h.locker.Lock()
	if ok,has := h.downloaded[key]; has && ok{
		ret = true
	}else{
		ret = false
	}
	h.locker.Unlock()
	return ret
}
