package cmd

import (
    "fmt"

    "github.com/jpedro/canned"
    "github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
    Use:   "rm",
    Short: "Removes an item",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        can, err := canned.OpenCan(CAN_FILE, CAN_PASSWORD)
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
