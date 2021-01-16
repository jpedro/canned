package cmd

import (
    "fmt"

    "github.com/jpedro/can"
    "github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
    Use:   "rm",
    Short: "Removes an item",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        store, err := can.OpenStore(CAN_FILE)
        if err != nil {
            panic(err)
        }

        err = store.DelItem(name)
        if err != nil {
            panic(err)
        }

        err = store.Save()
        if err != nil {
            panic(err)
        }

        fmt.Printf("==> Item %s removed.\n", paint("green", name))
    },
}

func init() {
    rootCmd.AddCommand(rmCmd)
}
