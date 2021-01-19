package main

import (
	"fmt"
	"os"
)

func main() {
	//创建一个名为“test”的目录，perm权限为0777
	err := os.Mkdir("test", 0777)
	if err != nil {
		fmt.Println(err)
	}
}