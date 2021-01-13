package cmd

import (
    "fmt"
    "../lib"

    "github.com/spf13/cobra"
    "github.com/atotto/clipboard"
)

var getCmd = &cobra.Command{
    Use:   "get",
    Short: "Get an item",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        can, err := lib.Open(CAN_FILE)
        if err != nil {
            panic(err)
        }

        item, err := can.GetItem(name)
        if err != nil {
            panic(err)
        }

        clipboard.WriteAll(item.Value)
        fmt.Printf("==> Item %s copied to the clipboard.\n", paint("green", name))
    },
}

func init() {
    rootCmd.AddCommand(getCmd)
}
