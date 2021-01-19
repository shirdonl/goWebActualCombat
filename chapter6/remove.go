package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("dir1")
	if err != nil {
		log.Fatal(err)
	}
}