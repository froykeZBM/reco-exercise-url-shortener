package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"reco-exercise-url-shortener/base62"
	redirector "reco-exercise-url-shortener/redirect"
	"reco-exercise-url-shortener/storage"
	"reco-exercise-url-shortener/url_generator"
)

func redirectLongUrl(w http.ResponseWriter, r *http.Request) {
	longUrl, err := redirector.GetOriginalUrl(r.URL.String()[1:])
	var status int
	switch err {
	case nil:
		status = http.StatusFound
	case storage.NotFoundInDB:
		status = http.StatusNotFound
	default:
		status = http.StatusBadGateway
	}
	http.Redirect(w, r, longUrl, status)
}

func storeUrlAndReturnShort(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	longUrl := string(reqBody)
	id := url_generator.CreateID(longUrl)
	fmt.Println("adding new url with id", longUrl, id)
	err = storage.AddUrl(longUrl, id)
	if err != nil {
		_, err := w.Write([]byte("ID is in use"))
		if err != nil {
			fmt.Println("Failed to send back the shortened url")
			return
		}
	}
	shortUrl := base62.Encode(id)
	shortUrlB := []byte(shortUrl)
	_, err = w.Write(shortUrlB)
	if err != nil {
		// For some reason we can't answer to teh request - just return to keep running.
		return
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/" {
	//	http.Error(w, "404 not found.", http.StatusNotFound)
	//	return
	//}

	switch r.Method {
	case "GET":
		redirectLongUrl(w, r)
	case "POST":
		storeUrlAndReturnShort(w, r)
	}
}
