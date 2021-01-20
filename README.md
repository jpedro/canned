# Canned
![Build Status](https://github.com/jpedro/canned/workflows/ci/badge.svg)
![Build Status](https://action-badges.now.sh/jpedro/canned)

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

Check [cli/canned](cli/canned) for its usage.
