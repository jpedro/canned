package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/jpedro/canned/lib"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a new item",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ensureFileExists()
		ensureWeHaveThePassword()

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
