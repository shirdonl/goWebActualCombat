package main

import "fmt"

func main() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 3:
				fmt.Println(i, j)
				continue OuterLoop //结束当前循环，开启下一次的外层循环
			}
		}
	}
}
