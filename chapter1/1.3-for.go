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
	//product := 1
	//for i := 1; i < 5; i++ {
	//	product *= i
	//}
	//println(product)
	//
	//	i := 0
	//	for {
	//		i++
	//		if i > 50 {
	//			break
	//		}
	//	}
	//	println(i)//51
	//
JumpLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if i > 2 {
				break JumpLoop
			}
			fmt.Println(i)
		}
	}

	//
	//
	//	j:= 2
	//	for ; j > 0; j-- {
	//		fmt.Println(j)
	//	}
	//
	//
	//	var i int
	//JumpLoop:
	//	for ; ; i++ {
	//		if i > 10 {
	//			//println(i)
	//			break JumpLoop
	//		}
	//	}

	//
	//	var i int
	//	for {
	//		if i > 10 {
	//			break
	//		}
	//		i++
	//	}
	//
	//	var i int
	//	for i <= 20 {
	//		println(i)
	//		i++
	//	}

}
