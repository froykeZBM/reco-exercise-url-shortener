package redirect

import (
	"fmt"
	"net/http"
	"net/url"
	"reco-exercise-url-shortener/base62"
	"reco-exercise-url-shortener/storage"
)

func shortenUrl(fullUrl string) (string, error) {

	url, err := url.Parse(fullUrl)

	if err != nil {
		return string(http.StatusBadRequest), fmt.Errorf("Invalid url")
	}

	host := url.Host

	id, err := base62.Decode(host)

	if err != nil {
		return "", err
	}

	longUrl := storage.GetUrl(id)
	url.Host = longUrl

	return url.String(), nil
}
