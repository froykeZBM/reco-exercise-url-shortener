package storage

import (
	"fmt"
	"time"
)

/*
 * Handle R/W actions on the url database
 */

type urlMapper map[uint64]item
type item struct {
	longUrl       string
	creationDate  time.Time
	LastVisitTime time.Time
}

var urlTable urlMapper
var NotFoundInDB error = fmt.Errorf("id not found")

/*
 * Init the database
 */
func initMapper() urlMapper {
	urlTable := make(urlMapper)
	return urlTable
}

/*
 * Get a long url from a short url in the db
 */
func GetUrl(id uint64) (string, error) {
	if !isInDB(id) {
		return "", NotFoundInDB
	}
	item := urlTable[id]
	// Update last request time:
	item.LastVisitTime = time.Now()
	urlTable[id] = item

	return item.longUrl, nil

}

/*
* add a pair of long and short url
 */
func AddUrl(newUrl string, id uint64) error {
	if isInDB(id) {
		return fmt.Errorf("id is taken")
	}
	newItem := item{newUrl, time.Now(), time.Now()}
	urlTable[id] = newItem
	return nil
}

func isInDB(id uint64) bool {
	_, ok := urlTable[id]
	if ok {
		return true
	} else {
		return false
	}
}
