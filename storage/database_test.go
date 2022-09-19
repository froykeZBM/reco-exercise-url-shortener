package storage

import "testing"

func TestDbClientInit(t *testing.T) {
	err := InitStorage()
	if err != nil {
		t.Errorf("Failed to init the db client")
	}

}

func TestSimpleGetSet(t *testing.T) {
	err := InitStorage()
	var exampleKey uint64 = 1234
	exampleUrl := "example.com"
	if err != nil {
		t.Errorf("Failed to init the db client")
	}
	err = AddUrl(exampleUrl, exampleKey)
	if err != nil {
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
