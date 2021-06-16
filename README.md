# Canned

[![Github Status](https://github.com/jpedro/canned/workflows/tests/badge.svg)](https://github.com/jpedro/canned/actions)
[![GoDoc](https://godoc.org/github.com/jpedro/canned?status.svg)](https://godoc.org/github.com/jpedro/canned)

Stores your secrets.


## Usage

```go
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
```

## CLI

Check [cli/canned](cli/canned) for your terminal needs.
