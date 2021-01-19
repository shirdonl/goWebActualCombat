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
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//var globalSession *SessionManager
//
//func init() {
//	globalSession, _ = NewSessionManager("memory", "GetSessionId", 3600)
//	//go globalSession.GarbageCollector()
//}

var globalSession *session.Manager
//然后在init函数中初始化
func init() {
	globalSession, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSession.GC()
}

type Session interface {
	Set(key, value interface{}) error //设置Session
	Get(key interface{}) interface{}  //获取Session
	Delete(key interface{}) error     //删除Session
	GetSessionId() string             //当前GetSessionId
}

type Provider interface {
	SessionInit(GetSessionId string) (Session, error)
	SessionRead(GetSessionId string) (Session, error)
	SessionDestroy(GetSessionId string) error
	GarbageCollector(maxLifeTime int64)
}

type SessionManager struct {
	cookieName  string     //cookie的名称
	lock        sync.Mutex //锁，保证并发时数据的安全一致
	provider    Provider   //管理session
	maxLifeTime int64      //超时时间
}

func RegisterProvider(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, p := providers[name]; p {
		panic("session: Register provider is existed")
	}
	providers[name] = provider
}

var providers = make(map[string]Provider)

func NewSessionManager(providerName, cookieName string, maxLifetime int64) (*SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}
	//返回一个 SessionManager 对象
	return &SessionManager{
		cookieName:  cookieName,
		maxLifeTime: maxLifetime,
		provider:    provider,
	}, nil
}

func (SessionManager *SessionManager) GetSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (SessionManager *SessionManager) SessionBegin(w http.ResponseWriter, r *http.Request) (session Session) {
	SessionManager.lock.Lock()
	defer SessionManager.lock.Unlock()
	cookie, err := r.Cookie(SessionManager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := SessionManager.GetSessionId()
		session, _ = SessionManager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: SessionManager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(SessionManager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = SessionManager.provider.SessionRead(sid)
	}
	return
}

//SessionDestroy 注销 Session
func (SessionManager *SessionManager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(SessionManager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}

	SessionManager.lock.Lock()
	defer SessionManager.lock.Unlock()

	SessionManager.provider.SessionDestroy(cookie.Value)
	expiredTime := time.Now()
	newCookie := http.Cookie{
		Name: SessionManager.cookieName,
		Path: "/", HttpOnly: true,
		Expires: expiredTime,
		MaxAge:  -1, //会话级cookie
	}
	http.SetCookie(w, &newCookie)
}

func (SessionManager *SessionManager) GarbageCollector() {
	SessionManager.lock.Lock()
	defer SessionManager.lock.Unlock()
	SessionManager.provider.GarbageCollector(SessionManager.maxLifeTime)
	//使用time包中的计时器功能，它会在session超时时自动调用GC方法
	time.AfterFunc(time.Duration(SessionManager.maxLifeTime), func() {
		SessionManager.GarbageCollector()
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	//if r.Method == "GET" {
	//	t, _ := template.ParseFiles("login.html")
	//	t.Execute(w, nil)
	//} else {
	//	sess := globalSession.SessionBegin(w, r)
	//	r.ParseForm()
	//	name := sess.Get("phone")
	//	if name != nil {
	//		sess.Set("phone", r.Form["phone"]) //将表单提交的phone值设置到session中
	//	}
	//}

	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, "aaaa")
	} else {
		sess := globalSession.SessionStart(w, r)
		sess.Set("phone", r.Form["phone"])
		//http.Redirect(w, r, "/", 302)
	}
}

func main() {
	http.HandleFunc("/login", login)         //设置路由
	err := http.ListenAndServe(":8088", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
