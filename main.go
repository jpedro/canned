package can

import (
    "time"
)

const (
    CAN_VERSION   = "v1"
    CAN_ALGORITHM = "GCM"
)

func NewStore(file string) (*Store, error) {
    store := &Store{}
    store.File = file
    store.Version = CAN_VERSION
    store.Metadata.CreatedAt = time.Now()
    store.Items = make(map[string]Item)

    return store, nil
}

func NewItem(name string, content string) (*Item, error) {
    item := &Item{}
    item.Content = content
    item.Metadata.CreatedAt = time.Now()
    item.Tags = []string{}

    return item, nil
}

func OpenStore(file string) (*Store, error) {
    store := &Store{}
    store.File = file
    err := store.Load()

    return store, err
}

func InitStore(file string) (*Store, error) {
    store, err := NewStore(file)
    if err != nil {
        return nil, err
    }

    err = store.Save()

    return store, err
}
