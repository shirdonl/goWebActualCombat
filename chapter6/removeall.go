package main

import (
	"log"
	"os"
)

func main() {
	//先创建多级子目录
	os.MkdirAll("test1/test2/test3", 0777)
	//删除test1目录及其子目录
	err := os.RemoveAll("test1")
	if err != nil {
		log.Fatal(err)
	}
}