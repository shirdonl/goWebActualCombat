package main

import (
	"fmt"
	"os"
)

func main() {
	//根据path创建多级子目录，例如dir1/dir2/dir3
	err :=os.MkdirAll("dir1/dir2/dir3", 0777)
	if err != nil {
		fmt.Println(err)
	}
}