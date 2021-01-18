package cmd

import (
	"fmt"

	"github.com/jpedro/canned"
	"github.com/spf13/cobra"
)

var tagRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes TAG from NAME",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ensureFile()
		ensurePassword()
		tag := args[0]
		name := args[1]

		can, err := canned.OpenCan(canFile, canPassword)
		if err != nil {
			bail("%", err)
		}

		ok := can.DelTag(name, tag)
		if !ok {
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
