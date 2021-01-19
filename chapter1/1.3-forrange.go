//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

import "fmt"

func main() {

	//for key, value := range []int{1, 2, 3, 4} {
	//	fmt.Printf("key:%d  value:%d\n", key, value)
	//}

	var str = "hi 加油"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	//m := map[string]int{
	//	"go": 100,
	//	"web": 100,
	//}
	//for key, value := range m {
	//	fmt.Println(key, value)
	//}

	c := make(chan int) //创建了一个整型类型的通道
	go func() {         //启动了一个协程
		c <- 7 //逻辑的实现
		c <- 8
		c <- 9
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}

	//for key, value := range []int{0, 1, -1, -2} {
	//	fmt.Printf("key:%d  value:%d\n", key, value)
	//}

	//m := map[string]int{
	//	"shirdon": 100,
	//	"ronger": 98,
	//}
	//for _, value := range m {
	//	fmt.Println(value)
	//}

	//for key, _ := range []int{9, 8, 7, 6} {
	//	fmt.Printf("key:%d \n", key)
	//}
}
