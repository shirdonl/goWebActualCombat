package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I_Love_Go_Web"
	res := strings.Split(s, "_")
	for i := range res {
		fmt.Println(res[i])
	}
	res1 := strings.SplitN(s, "_", 2)
	for i := range res1 {
		fmt.Println(res1[i])
	}
	res2 := strings.SplitAfter(s, "_")
	for i := range res2 {
		fmt.Println(res2[i])
	}
	res3 := strings.SplitAfterN(s, "_", 2)
	for i := range res3 {
		fmt.Println(res3[i])
	}
}