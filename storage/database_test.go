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
	newUrlItem.longUrl = "hi there"
	err = json.Unmarshal(data, &newUrlItem)
	if err != nil {
		t.Errorf("%v", err)
	}
	if newUrlItem != origUrlItem {
		t.Errorf("mismatch in marshalling: %v vs %v", newUrlItem, origUrlItem)
	}
}
