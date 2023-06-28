package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jpedro/canned/lib"
)

var tagAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds TAG to NAME",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ensureFileExists()
		ensureWeHaveThePassword()
		name := args[0]
		tag := args[1]

		can, err := canned.OpenCan(canFile, canPassword)
		if err != nil {
			bail("%s.", err)
		}

		err = can.AddTag(name, tag)
		if err != nil {
			bail("Tag %s was not added to item %s: %s\n",
				paint("green", tag),
				paint("green", name), err)
		}

		err = can.Save()
		if err != nil {
			bail("Failed to save file %s: %s\n", paint("green", canFile), err)
		}

		fmt.Printf("Tag %s was added to item %s\n",
			paint("green", tag),
			paint("green", name))
	},
}

func init() {
	tagCmd.AddCommand(tagAddCmd)
}
