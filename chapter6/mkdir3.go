package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
	err := os.MkdirAll(uploadDir , 777)
	if err!=nil{
		fmt.Println(err)
	}
}
