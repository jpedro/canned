package canned

import (
	"fmt"
	"time"
)

const (
	VERSION   = "v1"
	ALGORITHM = "GCM"
	SEPARATOR = "\n\n"
	DB_EXT    = ".sqlite"
)

// Creates a new can object
func NewCan(file string, password string) (*Can, error) {
	err := ensurePasswordExists(password)
	if err != nil {
		return nil, err
	}

	can := &Can{}
	can.file = file
	can.password = password
	can.Version = VERSION
	can.Algorithm = ALGORITHM
	can.Metadata.Created = time.Now()
	can.Items = make(map[string]*Item)

	return can, nil
}

// Creates a new item object
func NewItem(name string, content string) (*Item, error) {
	if name == "" {
		return nil, fmt.Errorf("name can't be empty")
	}

	item := &Item{}
	item.Content = content
	item.Metadata.Created = time.Now()

	return item, nil
}

// Opens a can file
// If the file doesn't exist, it will fail to load
func OpenCan(file string, password string) (*Can, error) {
	can, err := NewCan(file, password)
	if err != nil {
		return nil, err
	}

	err = can.load()
	if err != nil {
		return nil, err
	}

	err = openDatabase(file+DB_EXT, password)
	if err != nil {
		panic(err)
	}

	return can, err
}

// Initializes a new can file and will overwrite any existing file.
func InitCan(file string, password string) (*Can, error) {
	can, err := NewCan(file, password)
	if err != nil {
		return nil, err
	}

	err = initDatabase(file+DB_EXT, password)
	if err != nil {
		return nil, err
	}

	err = can.Save()
	if err != nil {
		return nil, err
	}

	return can, err
}
