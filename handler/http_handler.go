package handler

import (
	"net/http"
	redirector "reco-exercise-url-shortener/redirect"
	"reco-exercise-url-shortener/storage"
	"reco-exercise-url-shortener/url_generator"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		longUrl, err := redirector.GetOriginalUrl(r.URL.String())
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
	case "POST":
		longUrl := r.URL.String()
		id := url_generator.CreateID(longUrl)
		err := storage.AddUrl(longUrl, id)
		if err != nil {
			_, err := w.Write([]byte("ID is in use"))
			if err != nil {
				return
			}
		}
	}
}
