package main

import (
	"log"
)

func main() {
	no := []int{6, 8}
	log.Print("Print NO. ", no, "\n")
	log.Println("Println NO.", no)
	log.Printf("Printf NO. with item [%d,%d]\n", no[0], no[1])
}
