package main

import (
    "fmt"

    "github.com/jpedro/canned"
)

func main() {
    can, err := canned.InitCan("example.can")
    if err != nil {
        panic(err)
    }

    name  := "hello"
    value := "world"

    err = can.SetItem(name, value)
    if err != nil {
        panic(err)
    }
    err = can.Save()
    if err != nil {
        panic(err)
    }

    item, err := can.GetItem(name)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Item content: %s\n", item.Content)
}
