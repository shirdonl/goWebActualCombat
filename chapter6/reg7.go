package main

import (
	"fmt"
	"regexp"
)

func main() {
	res := findEmail("8888@qq.com")
	fmt.Println(res) // true

	res = findEmail("shir?don@qq.com")
	fmt.Println(res) // false

	res = findEmail("8888@qqcom")
	fmt.Println(res) // false
}
func findEmail(str string) bool {
	reg := regexp.MustCompile("^[a-zA-Z0-9_]+@[a-zA-Z0-9]+\\.[a-zA-Z0-9]+")
	res := reg.FindAllString(str, -1)
	if (res == nil) {
		return false
	}
	return true
}
