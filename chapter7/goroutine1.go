package main

import (
	"fmt"
	"time"
)

func Echo(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go Echo("go")
	Echo("web program")
}