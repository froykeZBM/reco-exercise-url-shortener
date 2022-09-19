package main

import (
	"log"
	"net/http"
	"reco-exercise-url-shortener/handler"
	"reco-exercise-url-shortener/storage"
)

func main() {
	err := storage.InitStorage()
	if err != nil {
		// A critical error of initializing the client happaned, so quit
		// TODO: better logging of said error
		log.Fatal(err)
		return
	}
	http.HandleFunc("/", handler.HandleRequest)

	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
