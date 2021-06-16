package cmd

import (
	"fmt"

	"github.com/jpedro/canned"
	"github.com/spf13/cobra"
)

var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Renames an item",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ensureFile()
		ensurePassword()
		name := args[0]
		new := args[1]

		if name == new {
			bail("New name is the same as the current one.")
		}

		can, err := canned.OpenCan(canFile, canPassword)
		if err != nil {
			bail("%s", err)
		}

		err = can.RenameItem(name, new)
		if err != nil {
			bail("%s", err)
		}

		err = can.Save()
		if err != nil {
			bail("%s", err)
		}

		fmt.Printf("Item %s renamed to %s.\n",
			paint("green", name),
			paint("green", new))
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)
}
