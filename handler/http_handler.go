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

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/" {
	//	http.Error(w, "404 not found.", http.StatusNotFound)
	//	return
	//}

	switch r.Method {
	case "GET":
		fmt.Println(r.URL.String())
		longUrl, err := redirector.GetOriginalUrl(r.URL.String()[1:])
		fmt.Println("long url is:", longUrl)
		var status int
		switch err {
		case nil:
			status = http.StatusFound
		case storage.NotFoundInDB:
			status = http.StatusNotFound
		default:
			status = http.StatusBadGateway
		}
		fmt.Println(longUrl)
		http.Redirect(w, r, longUrl, status)
	case "POST":
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
		w.Write(shortUrlB)
	}
}
