package cmd

import (
    "fmt"

    "github.com/jpedro/can"
    "github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Initializes a new can file",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Will initialize file %v.\n", paint("green", CAN_FILE))
        _, err := can.InitStore(CAN_FILE)
        if err != nil {
            panic(err)
        }
    },
}

func init() {
    rootCmd.AddCommand(initCmd)
}
