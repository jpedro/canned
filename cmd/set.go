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
		fmt.Printf("==> Setting %s: %s\n", name, value)
		can, err := lib.Open(CAN_FILE)
		if err != nil {
			panic(err)
		}

		can.SetItem(name, value)
		can.Save()
    },
}

func init() {
    rootCmd.AddCommand(setCmd)
}
