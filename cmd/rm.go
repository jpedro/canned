package cmd

import (
    "fmt"
    "../lib"

    "github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
    Use:   "rm",
    Short: "Removes an item",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        can, err := lib.Open(CAN_FILE)
        if err != nil {
            panic(err)
        }

        err = can.DelItem(name)
        if err != nil {
            panic(err)
		}

        err = can.Save()
        if err != nil {
            panic(err)
		}

        fmt.Printf("==> Item %s removed.\n", paint("green", name))
    },
}

func init() {
    rootCmd.AddCommand(rmCmd)
}
