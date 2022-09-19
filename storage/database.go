package storage

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
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

const redisAddress string = "172.22.112.1:6379"

var client *redis.Client

func InitStorage() error {
	UrlTable = make(urlMapper, 10000)
	//return UrlTable
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	return err
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
