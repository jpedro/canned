package cmd

import (
    "fmt"

    "../lib"

    "github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Initializes a new can file",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Will initialize file %v.\n", paint("green", CAN_FILE))
        lib.Init(CAN_FILE)
    },
}

func init() {
    rootCmd.AddCommand(initCmd)
}
