package main

import (
	"fmt"

	"github.com/jpedro/canned"
)

func main() {
	file := "/tmp/example.can"
	password := "test123"
	name := "hello"
	value := "world"

	can, _ := canned.InitCan(file, password)
	can.SetItem(name, value)
	can.Save()

	can, _ = canned.OpenCan(file, password)
	item, _ := can.GetItem(name)

	fmt.Printf("Item '%s' content: '%s'.\n", name, item.Content)
}
