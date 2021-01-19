package main

import (
	"log"
	"os"
)

func main() {
	fileName := "New.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	debugLog := log.New(logFile, "[Info]", log.Llongfile)
	debugLog.Println("Info Level Message")
	debugLog.SetPrefix("[Debug]")
	debugLog.Println("Debug Level Message")
}
