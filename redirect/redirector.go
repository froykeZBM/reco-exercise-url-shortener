package redirector

import (
	"errors"
	"fmt"
	"net/url"
	"reco-exercise-url-shortener/base62"
	"reco-exercise-url-shortener/storage"
)

var InvalidEncoding = errors.New("invalid encoding")
var URLParseError = fmt.Errorf("invalid parsedUrl")

func GetOriginalUrl(fullUrl string) (string, error) {

	parsedUrl, err := url.Parse(fullUrl)
	parsedUrl.Query()
	if err != nil {
		return "", URLParseError
	}

	host := parsedUrl.String()
	fmt.Println("host is:", host)
	id, err := base62.Decode(host)
	fmt.Println("Found ID:", id)
	if err != nil {
		return "", InvalidEncoding
	}

	longUrl, err := storage.GetUrl(id)

	if err != nil {
		return "", err
	}
	return longUrl, nil
}
