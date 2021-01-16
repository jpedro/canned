package canned

import (
    "time"
)

const (
    VERSION   = "v1"
    ALGORITHM = "GCM"
    SEPARATOR = "\n\n"
)

// Creates a new can in memory
func NewCan(file string, password string) (*Can, error) {
    if err := verifyPassword(password); err != nil {
        return nil, err
    }

    can := &Can{}
    can.File = file
    can.Password = password
    can.Version = VERSION
    can.Metadata.CreatedAt = time.Now()
    can.Items = make(map[string]Item)

    return can, nil
}

// Creates a new item in memory
func NewItem(name string, content string) (*Item, error) {
    item := &Item{}
    item.Content = content
    item.Metadata.CreatedAt = time.Now()
    // item.Tags = []string{}

    return item, nil
}

// Opens a can file
func OpenCan(file string, password string) (*Can, error) {
    can, err := NewCan(file, password)
    if err != nil {
        return nil, err
    }
    err = can.Load()

    return can, err
}

// Initializes a new can file
func InitCan(file string, password string) (*Can, error) {
    can, err := NewCan(file, password)
    if err != nil {
        return nil, err
    }

    err = can.Save()

    return can, err
}
