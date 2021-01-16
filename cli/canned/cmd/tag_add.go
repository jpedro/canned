package cmd

import (
    "fmt"

    "github.com/jpedro/canned"
    "github.com/spf13/cobra"
)

var tagAddCmd = &cobra.Command{
    Use:   "add",
    Short: "Adds TAG to NAME",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
		ensurePassword()
        tag  := args[0]
        name := args[1]

        can, err := canned.OpenCan(CAN_FILE, CAN_PASSWORD)
        if err != nil {
            panic(err)
        }

        err = can.AddTag(name, tag)
        if err != nil {
            bail("==> Tag %s was not added to %s.\n", paint("green", tag), paint("green", name))
        }

        err = can.Save()
        if err != nil {
            bail("==> Failed to save file %s%s.\n", paint("green", CAN_FILE))
        }

        fmt.Printf("==> Tag %s was added to %s\n", paint("green", tag), paint("green", name))
    },
}

func init() {
    tagCmd.AddCommand(tagAddCmd)
}
