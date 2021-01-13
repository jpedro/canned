package lib

import (
    "time"
)

type Item struct {
    Name        string      `json:"name"`
    Value       string      `json:"value"`
    Metadata    Metadata    `json:"metadata"`
    Tags        []string    `json:"tags"`
}

func NewItem(name string, value string) (*Item, error) {
    item := &Item{}
    item.Name = name
    item.Value = value
    item.Metadata.CreatedAt = time.Now()
    // item.UpdatedAt = time.Time{}
    item.Tags = []string{}

    return item, nil
}
