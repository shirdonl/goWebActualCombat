package main

import "fmt"

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["a"] = "one" // 第一个冲突访问.
		c <- true
	}()
	m["b"] = "two" // 第一个冲突访问
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
