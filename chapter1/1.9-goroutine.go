package main

import (
	"fmt"
	"time"
)

func HelloWorld() {
	fmt.Println("this is a goroutine msg")
}

func main() {
	go HelloWorld()
	time.Sleep(1 * time.Second)
	fmt.Println("end")
}
