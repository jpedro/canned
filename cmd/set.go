package cmd

import (
    "fmt"
    "../lib"

    "github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
    Use:   "set",
    Short: "Sets a new item",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        value := args[1]

		can, err := lib.Open(CAN_FILE)
        if err != nil {
            panic(err)
        }

        can.SetItem(name, value)
        can.Save()

		fmt.Printf("==> Item %s saved.\n", paint("green", name))
    },
}

func init() {
    rootCmd.AddCommand(setCmd)
}
