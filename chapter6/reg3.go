package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I_Love_Go_Web"
	res := strings.Split(s, "_")
	for value := range res {
		fmt.Println(value)
	}
}