package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jpedro/canned/lib"
)

var tagRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes TAG from NAME",
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

		err = can.DelTag(name, tag)
		if err != nil {
			bail("Tag %s was not removed from %s.\n",
				paint("green", tag),
				paint("green", name))
		}

		err = can.Save()
		if err != nil {
			bail("Failed to save file %s.\n",
				paint("green", canFile))
		}

		fmt.Printf("Tag %s removed from %s\n",
			paint("green", tag),
			paint("green", name))
	},
}

func init() {
	tagCmd.AddCommand(tagRmCmd)
}
