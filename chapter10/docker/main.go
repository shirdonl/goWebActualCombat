package main

import (
	"fmt"
	"log"
	"net/http"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This server is built by Docker!")
}

func main() {
	http.HandleFunc("/", hi)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}