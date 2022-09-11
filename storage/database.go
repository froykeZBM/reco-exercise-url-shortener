package storage

import (
	"time"
)

/*
 * Handle R/W actions on the url database
 */

type urlMapper map[uint64]item
type item struct {
	id            uint64
	longUrl       string
	creationDate  time.Time
	LastVisitDate time.Time
}

var urlTable urlMapper

/*
 * Init the database
 */
func initMapper(url main.shortUrl) urlMapper {
	urlTable := make(urlMapper)
	return urlTable
}

/*
 * Get a long url from a short url in the db
 */
func GetUrl(id uint64) string {

}

/*
* add a pair of long and short url
 */
func addUrl(short main.shortUrl, long longUrl) {
	_, ok := urlTable[short]
	if !ok {
		urlTable[short] = long
	}
}
