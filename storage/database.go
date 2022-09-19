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
	LongUrl          string
	LastCreationDate time.Time
	LastVisitTime    time.Time
	VisitNum         int32
	CreationNum      int32
}

var UrlTable urlMapper
var NotFoundInDB = fmt.Errorf("id not found")
var FatalDbError = fmt.Errorf("fatal error in database handling")
var SaveError = fmt.Errorf("Could not save in db.")
var IdTakenError = fmt.Errorf("Id is taken in the DB")

// I'm having trouble configuring the redis container.
// This is the IP address of the container for now (and it works)
const redisAddress string = "172.22.112.1:6379"

var client *redis.Client

// A function to initialize the DB client to interact with it
// Might be better to call this every time an access is required to the server
func InitStorage() error {
	UrlTable = make(urlMapper, 10000)
	//return UrlTable
	client = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
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
	return item.LongUrl, errReturned

}

func createNewUrl(newUrl string) urlItem {
	return urlItem{newUrl, time.Now(), time.Now(), 0, 1}
}

func updateExistingUrl(data_str string) (urlItem, error) {
	var item urlItem
	err := json.Unmarshal([]byte(data_str), &item)
	if err != nil {
		return urlItem{}, err
	}
	item.CreationNum += 1
	item.LastCreationDate = time.Now()
	return item, nil

}

func AddUrl(newUrl string, id uint64) error {
	id_str := fmt.Sprint(id)
	data_str, err := client.Get(id_str).Result()
	var newUrlItem urlItem
	if err == redis.Nil {
		newUrlItem = createNewUrl(newUrl)
	} else if err != nil {
		return err
	} else {
		newUrlItem, err = updateExistingUrl(data_str)
		if err != nil {
			return err
		}
	}

	data, err := json.Marshal(newUrlItem)
	if err != nil {
		fmt.Println(err)
		return SaveError
	}
	_, err = client.Set(id_str, data, 0).Result()
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
