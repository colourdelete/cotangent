package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("will listen on :9090")
	log.Println("only for testing")
	err := http.ListenAndServe(":9090", http.FileServer(http.Dir(".")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
