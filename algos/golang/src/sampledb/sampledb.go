package main

import (
	"fmt"
	"time"

	"github.com/timshannon/bolthold"
)

const filename = "sample.db"

// Item structure
type Item struct {
	Name    string
	Created time.Time
}

func main() {
	store, err := bolthold.Open(filename, 0666, nil)
	if err != nil {
		//handle error
	}
	err = store.Insert("key", &Item{
		Name:    "Test Name",
		Created: time.Now(),
	})

	err = store.Insert("key1", &Item{
		Name:    "Test Name1",
		Created: time.Now(),
	})

	result := &Item{}
	fmt.Println(store.Get("key1", result))
	fmt.Println("Result", result)

}
