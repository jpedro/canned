package main

import (
	"fmt"

	"github.com/jpedro/canned"
)

func main() {
	password := "test"
	file := "example.can"

	can, err := canned.InitCan(file, password)
	if err != nil {
		panic(err)
	}

	err = can.SetItem("hello", "world")
	if err != nil {
		panic(err)
	}

	err = can.Save()
	if err != nil {
		panic(err)
	}

	can2, err := canned.OpenCan(file, password)
	if err != nil {
		panic(err)
	}

	item, err := can2.GetItem("hello")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Item content: %s\n", item.Content)
}
