package storage

import "testing"

func TestDbClient(t *testing.T) {
	err := InitStorage()
	if err != nil {
		t.Errorf("Failed to init the db client")
	}

}
