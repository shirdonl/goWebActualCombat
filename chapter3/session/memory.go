package main

import (
	"container/list"
	"sync"
	"time"
)

var provider = &FromMemory{list: list.New()}

func init() {
	provider.sessions = make(map[string]*list.Element, 0)
	//注册  memory 调用的时候一定要一致
	RegisterProvider("memory", provider)
}

//session实现
type SessionStore struct {
	sid              string                      //session id 唯一标示
	LastAccessedTime time.Time                   //最后访问时间
	value            map[interface{}]interface{} //session 里面存储的值
}

//设置
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	provider.SessionUpdate(st.sid)
	return nil
}

//获取session
func (st *SessionStore) Get(key interface{}) interface{} {
	provider.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}

//删除
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	provider.SessionUpdate(st.sid)
	return nil
}
func (st *SessionStore) SessionID() string {
	return st.sid
}

//session来自内存 实现
type FromMemory struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做 gc
}

func (frommemory *FromMemory) SessionInit(sid string) (Session, error) {
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, LastAccessedTime: time.Now(), value: v}
	element := frommemory.list.PushBack(newsess)
	frommemory.sessions[sid] = element
	return newsess, nil
}

func (frommemory *FromMemory) SessionRead(sid string) (Session, error) {
	if element, ok := frommemory.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := frommemory.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (frommemory *FromMemory) SessionDestroy(sid string) error {
	if element, ok := frommemory.sessions[sid]; ok {
		delete(frommemory.sessions, sid)
		frommemory.list.Remove(element)
		return nil
	}
	return nil
}

func (frommemory *FromMemory) SessionGC(maxLifeTime int64) {
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	for {
		element := frommemory.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).LastAccessedTime.Unix() + maxLifeTime) <
			time.Now().Unix() {
			frommemory.list.Remove(element)
			delete(frommemory.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}
func (frommemory *FromMemory) SessionUpdate(sid string) error {
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	if element, ok := frommemory.sessions[sid]; ok {
		element.Value.(*SessionStore).LastAccessedTime = time.Now()
		frommemory.list.MoveToFront(element)
		return nil
	}
	return nil
}
