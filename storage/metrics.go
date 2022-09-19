package storage

import (
	"encoding/json"
	"fmt"
	"time"
)

// This is a function to get the metadata of the url
// Next step is to add a handler to the http server ot get the metrics
func GetUrlStruct(id uint64) (urlItem, error) {
	id_str := fmt.Sprint(id)
	data, err := client.Get(id_str).Result()
	if err != nil {
		return urlItem{}, err
	}

	// I'm not sure if using the json encoding is neccessary
	// Pros: simpler storage in the db and easy serilizing
	// Cons: might increase a single entry size in the db.
	var item urlItem
	err = json.Unmarshal([]byte(data), &item)
	if err != nil {
		return urlItem{}, err
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
	return item, errReturned
}
