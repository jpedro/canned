package cmd

import (
	"fmt"

	"github.com/jpedro/canned"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a new item",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ensureFile()
		ensurePassword()
		name := args[0]
		value := args[1]

		can, err := canned.OpenCan(canFile, canPassword)
		if err != nil {
			bail("%s.", err)
		}

		err = can.SetItem(name, value)
		if err != nil {
			bail("%s.", err)
		}

		err = can.Save()
		if err != nil {
			bail("%s.", err)
		}

		fmt.Printf("Item %s stored.\n", paint("green", name))
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
