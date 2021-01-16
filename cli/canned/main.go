package main

import (
    "os"

    "github.com/jpedro/can/cli/can/cmd"
)

func main() {
    err := cmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}
