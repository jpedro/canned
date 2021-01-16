package cmd

import (
    "fmt"

    "github.com/jpedro/can"
    "github.com/spf13/cobra"
)

var tagAddCmd = &cobra.Command{
    Use:   "add",
    Short: "Adds TAG to NAME",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        tag  := args[0]
        name := args[1]

        store, err := can.OpenStore(CAN_FILE)
        if err != nil {
            panic(err)
        }

        err = store.AddTag(name, tag)
        if err != nil {
            bail("==> Tag %s was not added to %s.\n", paint("green", tag), paint("green", name))
        }

        err = store.Save()
        if err != nil {
            bail("==> Failed to save file %s%s.\n", paint("green", CAN_FILE))
        }

        fmt.Printf("==> Tag %s was added to %s\n", paint("green", tag), paint("green", name))
    },
}

func init() {
    tagCmd.AddCommand(tagAddCmd)
}
