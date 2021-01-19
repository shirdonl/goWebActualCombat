package main

import (
	"fmt"
	"strings"
)

func main() {
	value := "a|b|c|d"
	// 分割成3部分
	result := strings.SplitN(value, "|", 3)
	for v := range(result) {
		fmt.Println(result[v])
	}
}