package redirector

import (
	"errors"
	"fmt"
	"net/url"
	"reco-exercise-url-shortener/base62"
	"reco-exercise-url-shortener/storage"
)

var InvalidEncoding error = errors.New("invalid encoding")
var URLParseError error = fmt.Errorf("invalid parsedUrl")

func GetOriginalUrl(fullUrl string) (string, error) {

	parsedUrl, err := url.Parse(fullUrl)

	if err != nil {
		return "", URLParseError
	}

	host := parsedUrl.Host

	id, err := base62.Decode(host)

	if err != nil {
		return "", InvalidEncoding
	}

	longUrl, err := storage.GetUrl(id)

	if err != nil {
		return "", err
	}
	parsedUrl.Host = longUrl

	return parsedUrl.String(), nil
}
