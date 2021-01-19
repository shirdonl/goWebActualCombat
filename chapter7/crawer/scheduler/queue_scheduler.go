//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package scheduler

import (
	"container/list"
	"crypto/md5"
	"sync"
)

type QueueScheduler struct {
	queue    *list.List
	locker   *sync.Mutex
	listkey  map[[md5.Size]byte] *list.Element
}

func NewQueueSCheduler() *QueueScheduler{
	queue   := list.New()
	locker  := new(sync.Mutex)
	listkey := make(map[[md5.Size]byte] *list.Element)

	return &QueueScheduler{
		queue:queue,
		locker:locker,
		listkey:listkey}
}

// Pop - 从队列中获取一个链接
func (s *QueueScheduler) Pop() (string,bool){
	s.locker.Lock()
	if s.queue.Len() <= 0{
		s.locker.Unlock()
		return "",false
	}
	e := s.queue.Front()
	ret := e.Value.(string)
	// 清除listkey中该元素,加入到访问队列中
	key := md5.Sum([]byte(ret))
	delete(s.listkey,key)
	s.queue.Remove(e)
	s.locker.Unlock()
	return ret,true
}

// Push - 将链接放入队列中
func (s *QueueScheduler) Push(url string){
	s.locker.Lock()
	key := md5.Sum([]byte(url))
	// 链接已存在
	if _,ok := s.listkey[key]; ok{
		s.locker.Unlock()
		return
	}
	e := s.queue.PushBack(url)
	s.listkey[key] = e
	s.locker.Unlock()
}
