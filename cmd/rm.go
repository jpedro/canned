package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jpedro/canned/lib"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes an item",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ensureFileExists()
		ensureWeHaveThePassword()

		name := args[0]
		can, err := canned.OpenCan(canFile, canPassword)
		if err != nil {
			bail("%s.", err)
		}

		err = can.DelItem(name)
		if err != nil {
			bail("%s.", err)
		}

		err = can.Save()
		if err != nil {
			bail("%s.", err)
		}

		fmt.Printf("Item '%s' removed.\n", paint("green", name))
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
