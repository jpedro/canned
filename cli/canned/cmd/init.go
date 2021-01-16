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
        _, err := can.InitCan(CAN_FILE, CAN_PASSWORD)
        if err != nil {
            panic(err)
        }
    },
}

func init() {
    rootCmd.AddCommand(initCmd)
}
