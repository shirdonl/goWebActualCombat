package main

import "fmt"

func main() {
OuterLoop: //外层循环的标签
	for i := 0; i < 2; i++ { //双层循环
		for j := 0; j < 5; j++ {
			switch j { //使用 switch 进行数值分支判断
			case 1:
				fmt.Println(i, j)
				break OuterLoop
			case 2:
				fmt.Println(i, j)
				break OuterLoop //退出 OuterLoop 对应的循环之外
			}
		}
	}
}
