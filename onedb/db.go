package onedb

import (
	"fmt"
)

var doc = make(map[int]string)

// Save adds a new item in the database.
// If there is an item with the given key
// returns an error.
func Save(key int, value string) error {

	if _, exist := doc[key]; exist {
		return fmt.Errorf("The given key [%d] already exist", key)
	}

	doc[key] = value

	return nil
}

// Find searches an item with the given key
// if the key doesn't exist, it returns an error
func Find(key int) (string, error) {

	value, exist := doc[key]

	if !exist {
		return "", fmt.Errorf("The key [%d] was not find", key)
	}

	return value, nil
}
