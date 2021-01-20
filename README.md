# Canned
[![Build Status](https://action-badges.now.sh/jpedro/canned)](https://github.com/jpedro/canned/actions)
[![Github Status](https://github.com/jpedro/canned/workflows/main/badge.svg)](https://github.com/jpedro/canned/actions)
[![GoDoc](https://godoc.org/github.com/jpedro/canned?status.svg)](https://godoc.org/github.com/jpedro/canned)

Go library to store encrypted goods.


## Usage

```go
package main

import (
    "fmt"

    "github.com/jpedro/canned"
)

func main() {
    can, _ := canned.InitCan("example.can")
    can.SetItem("name", "value")
    can.Save()
    item, _ := can.GetItem(name)
    fmt.Printf("Item content: %s\n", item.Content)
}
```

## CLI

Check [cli/canned](cli/canned) for your terminal needs.
