package cmd

import (
    "fmt"

    "github.com/jpedro/can"
    "github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
    Use:   "set",
    Short: "Sets a new item",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        value := args[1]

        store, err := can.OpenStore(CAN_FILE)
        if err != nil {
            panic(err)
        }

        err = store.SetItem(name, value)
        if err != nil {
            panic(err)
        }

        err = store.Save()
        if err != nil {
            panic(err)
        }

        fmt.Printf("==> Item %s saved.\n", paint("green", name))
    },
}

func init() {
    rootCmd.AddCommand(setCmd)
}
