package storage

import (
	"encoding/json"
	"testing"
	"time"
)

var exampleKey uint64 = 1234
var exampleUrl string = "example.com"

func TestDbClientInit(t *testing.T) {
	err := InitStorage()
	if err != nil {
		t.Errorf("Failed to init the db client")
	}

}

func TestSimpleGetSet(t *testing.T) {
	err := InitStorage()

	if err != nil {
		t.Errorf("Failed to init the db client")
	}
	err = AddUrl(exampleUrl, exampleKey)
	if (err != nil) && (err != IdTakenError) {
		t.Errorf("Failed with the following exception: %v", err)
	}
	url, err := GetUrl(exampleKey)
	if err != nil {
		t.Errorf("Failed with the following exception: %v", err)
	}
	if url != exampleUrl {
		t.Errorf("url mismatch - expected %v, but got %v", exampleUrl, url)
	}
}

func TestJsonEncoding(t *testing.T) {
	origUrlItem := urlItem{exampleUrl, time.Now(), time.Now(), 0, 1}
	data, err := json.Marshal(origUrlItem)
	if err != nil {
		t.Errorf("%v", err)
	}
	var newUrlItem urlItem
	err = json.Unmarshal(data, &newUrlItem)
	if err != nil {
		t.Errorf("%v", err)
	}
	if newUrlItem.LongUrl != origUrlItem.LongUrl {
		t.Errorf("mismatch in marshalling: %v vs %v", newUrlItem.LongUrl, origUrlItem.LongUrl)
	}
}

// TODO tests:
func TestUrlNotInDb(t *testing.T) {
	// How to make sure a url is not in the db? since this is persistent
	//  It's tricky. I might need to write a helper function to remove the url first
	// so, Steps:
	// 1. see if URL is in db
	// 3. 2. if yes, delete it
	// 4. try to get a url which is not in the DB
	// 5. expect a url not found exception
}

func TestUrlCreationIncreasing(t *testing.T) {
	// Steps:
	// 1. get a Url
	// 2. If no url exists, create one
	// 3. save it's number of creation attempts
	// 4. create the same url
	// 5. assert that the creation number has increased by 1
	// 6. asser that the time of last creation has changed
}

func TestUrlVisitIncreasing(t *testing.T) {
	// Steps:
	// 1. get(/create if not exists) a url
	// 2. get the url from the DB
	// 3. assert that the visitNum has increased
	// 4. assert changing in creation of time
}
