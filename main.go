package main

import (
	"fmt"
	"log"
	"net/http"
	"reco-exercise-url-shortener/handler"
	"reco-exercise-url-shortener/storage"
)

func main() {
	storage.InitMapper()
	http.HandleFunc("/", handler.HandleRequest)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
