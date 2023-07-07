package main

import (
    "github.com/jpedro/canned/cmd"
)

func main() {
    err := cmd.Execute()
    if err != nil {
        panic(err)
    }
}
