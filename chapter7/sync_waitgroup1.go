package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1 goroutine sleep ...")
		time.Sleep(2)
		fmt.Println("1 goroutine exit ...")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2 goroutine sleep ...")
		time.Sleep(4)
		fmt.Println("2 goroutine exit ...")
	}()

	fmt.Println("Waiting for all goroutine ")
	wg.Wait()
	fmt.Println("All goroutines finished!")
}
