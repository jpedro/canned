package cmd

import (
	"fmt"

	"github.com/jpedro/canned"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes an item",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ensureFile()
		ensurePassword()
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

		fmt.Printf("Item %s removed.\n", paint("green", name))
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
