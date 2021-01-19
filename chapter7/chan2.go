package main

import (
	"fmt"
)

func Sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum // 把 sum 发送到通道 ch
}

func main() {
	s := []int{6, 7, 8, -9, 1, 8}
	ch := make(chan int)
	go Sum(s[:len(s)/2], ch)
	go Sum(s[len(s)/2:], ch)
	a, b := <-ch, <-ch // 从通道 ch 中接收
	fmt.Println(a, b, a+b)
}
