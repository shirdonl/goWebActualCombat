package main

import "fmt"

func main() {
	//......省略前面代码
	err := getUserInfo()
	if err != nil {
		fmt.Println(err)
		exitProcess()
		return
	}
	err = getEmail()
	if err != nil {
		fmt.Println(err)
		exitProcess()
		return
	}
	fmt.Println("over")

//	err := getUserInfo()
//	if err != nil {
//		goto doExit //将跳转错误标签 onExit
//	}
//	err = getEmail()
//	if err != nil {
//		goto doExit //将跳转错误标签 onExit
//	}
//	fmt.Println("over")
//	return
//doExit: //汇总所有流程进行错误打印并退出进程
//	fmt.Println(err)
//	exitProcess()

}
