package handler

import (
	"errors"
	"fmt"
	"net/url"
	"reco-exercise-url-shortener/base62"
	"reco-exercise-url-shortener/storage"
)

func shortenUrl(fullUrl string) (string, error) {

	parsedUrl, err := url.Parse(fullUrl)

	if err != nil {
		return "", fmt.Errorf("invalid parsedUrl")
	}

	host := parsedUrl.Host

	id, err := base62.Decode(host)

	if err != nil {
		return "", errors.New("invalid encoding")
	}

	longUrl, err := storage.GetUrl(id)

	if err != nil {
		return "", err
	}
	parsedUrl.Host = longUrl

	return parsedUrl.String(), nil
}
