package cmd

import (
    "fmt"

    "github.com/jpedro/canned"
    "github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
    Use:   "set",
    Short: "Sets a new item",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        value := args[1]

        can, err := canned.OpenCan(CAN_FILE, CAN_PASSWORD)
        if err != nil {
            panic(err)
        }

        err = can.SetItem(name, value)
        if err != nil {
            panic(err)
        }

        err = can.Save()
        if err != nil {
            panic(err)
        }

        fmt.Printf("==> Item %s saved.\n", paint("green", name))
    },
}

func init() {
    rootCmd.AddCommand(setCmd)
}
