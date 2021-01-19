package main

import (
	"fmt"
	"strings"
)

func main() {
	value := "i'm gopher"
	result := strings.SplitAfter(value, "")
	for i := range (result) {
		// 获取字母
		letter := result[i]
		fmt.Println(letter)
	}
}
