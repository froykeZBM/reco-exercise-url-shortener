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

var UrlTable urlMapper
var NotFoundInDB = fmt.Errorf("id not found")

func InitMapper() /*urlMapper*/ {
	UrlTable = make(urlMapper, 10000)
	//return UrlTable
}

func GetUrl(id uint64) (string, error) {
	if !isInDB(id) {
		return "", NotFoundInDB
	}
	item := UrlTable[id]
	// Update last request time:
	item.LastVisitTime = time.Now()
	UrlTable[id] = item

	return item.longUrl, nil

}

func AddUrl(newUrl string, id uint64) error {
	if isInDB(id) {
		//TODO: check if new url matches the id, and then just update the creation time
		return fmt.Errorf("id is taken")
	}
	newItem := item{newUrl, time.Now(), time.Now()}
	UrlTable[id] = newItem
	return nil
}

func isInDB(id uint64) bool {
	_, ok := UrlTable[id]
	if ok {
		return true
	} else {
		return false
	}
}
