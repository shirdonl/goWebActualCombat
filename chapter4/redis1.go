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
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed, err:", err)
		return
	}
	defer c.Close()


	_, err = c.Do("Set", "username", "jim")
	if err != nil {
		fmt.Println(err)
		return
	}
	//
	//res, err := redis.String(c.Do("Get", "username"))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(res)
	//
	//_, err = c.Do("MSet", "username", "james", "phone", "18888888888")
	//if err != nil {
	//	fmt.Println("MSet error: ", err)
	//	return
	//}
	//
	//res2, err := redis.Strings(c.Do("MGet", "username", "phone"))
	//if err != nil {
	//	fmt.Println("MGet error: ", err)
	//	return
	//}
	//fmt.Println(res2)
	//
	//
	//
	//_, err = c.Do("HSet", "names", "jim", "barry")
	//if err != nil {
	//	fmt.Println("hset error: ", err)
	//	return
	//}
	//
	//res3, err := redis.String(c.Do("HGet", "names", "jim"))
	//if err != nil {
	//	fmt.Println("hget error: ", err)
	//	return
	//}
	//fmt.Println(res3)
	//
	//
	////expire
	//_, err = c.Do("expire", "names", 5)
	//if err != nil {
	//	fmt.Println("expire error: ", err)
	//	return
	//}


	//队列
	//_, err = c.Do("lpush", "Queue", "jim", "barry", 9)
	//if err != nil {
	//	fmt.Println("lpush error: ", err)
	//	return
	//}
	//for {
	//	r, err := redis.String(c.Do("lpop", "Queue"))
	//	if err != nil {
	//		fmt.Println("lpop error: ", err)
	//		break
	//	}
	//	fmt.Println(r)
	//}
	//res4, err := redis.Int(c.Do("llen", "Queue"))
	//if err != nil {
	//	fmt.Println("llen error: ", err)
	//	return
	//}
	//fmt.Println(res4)

}
