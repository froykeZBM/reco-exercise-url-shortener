package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

/*
 * Handle R/W actions on the url database
 */

type urlMapper map[uint64]urlItem
type urlItem struct {
	longUrl          string
	LastCreationDate time.Time
	LastVisitTime    time.Time
	VisitNum         int32
	CreationNum      int32
}

var UrlTable urlMapper
var NotFoundInDB = fmt.Errorf("id not found")
var FatalDbError = fmt.Errorf("fatal error in database handling")
var SaveError = fmt.Errorf("Could not save in db.")

const redisAddress string = "172.22.112.1:6379"

var client *redis.Client

// A function to initialize the DB client to interact with it
// Might be better to call this every time an access is required to the server
func InitStorage() error {
	UrlTable = make(urlMapper, 10000)
	//return UrlTable
	Client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})
	_, err := Client.Ping().Result()
	if err != nil {
		// if we cannot initialize the redis client, this is a fatal error
		return FatalDbError
	}
	return err
}

// Getting a url from the db.
// every request also updates the url in the DB
func GetUrl(id uint64) (string, error) {
	if !isInDB(id) {
		return "", NotFoundInDB
	}
	id_str := fmt.Sprint(id)
	data, err := client.Get(id_str).Result()
	if err != nil {
		return "", err
	}
	// I'm not sure if using the json encoding is neccessary
	// Pros: simpler storage in the db and easy serilizing
	// Cons: might increase a single entry size in the db.
	var item urlItem
	err = json.Unmarshal([]byte(data), &item)

	if err != nil {
		return "", err
	}
	// Update last request time:
	item.LastVisitTime = time.Now()
	item.VisitNum += 1
	newData, err := json.Marshal(item)

	var errReturned error = nil
	_, err = client.Set(id_str, newData, 0).Result()
	if err != nil {
		errReturned = SaveError
	}
	return item.longUrl, errReturned

}

func AddUrl(newUrl string, id uint64) error {
	if isInDB(id) {
		//TODO: check if new url matches the id, and then just update the creation time
		return fmt.Errorf("id is taken")
	}
	newurlItem := urlItem{newUrl, time.Now(), time.Now(), 0, 1}
	data, err := json.Marshal(newurlItem)
	if err != nil {
		fmt.Println(err)
		return SaveError
	}
	_, err = client.Set(fmt.Sprint(id), data, 0).Result()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func isInDB(id uint64) bool {
	_, err := client.Get(fmt.Sprint((id))).Result()
	if err == redis.Nil {
		return false
	} else {
		return true
	}
}
