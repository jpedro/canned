package main

import (
    "os"

    "github.com/jpedro/canned/cli/canned/cmd"
)

func main() {
    err := cmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}
