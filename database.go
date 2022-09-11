package reco_exercise_url_shortener

/*
 * Handle R/W actions on the url data base
 */
type longUrl string
type shortUrl string
type urlMapper map[shortUrl]longUrl

var urlTable urlMapper

/*
 * Init the database
 */
func initMapper(url shortUrl) urlMapper {
	urlTable := make(urlMapper)
	return urlTable
}

/*
 * Get a short url from a long url in the db
 */
func getShortFromLongUrl(url longUrl) {

}

/*
 * Get a long url from a short url in the db
 */
func getLongFromShortUrl(url longUrl) {

}

/*
* add a pair of long and short url
 */
func addUrl(short shortUrl, long longUrl) {
	_, ok := urlTable[short]
	if !ok {
		urlTable[short] = long
	}
}
