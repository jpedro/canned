package main

import (
	"fmt"

	"github.com/jpedro/canned"
)

func main() {
	can, _ := canned.InitCan("example.can", "pp")
	can.SetItem("hello", "world")
	can.Save()

	can2, _ := canned.OpenCan("example.can", "pp")
	item, _ := can2.GetItem("hello")
	fmt.Printf("Item content: %s\n", item.Content)
}
