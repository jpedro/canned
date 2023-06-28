package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"github.com/jpedro/canned/lib"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Copies an item to the clipboard",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ensureFileExists()
		ensureWeHaveThePassword()
		name := args[0]
		can, err := canned.OpenCan(canFile, canPassword)
		if err != nil {
			bail("%s.", err)
		}

		item, err := can.GetItem(name)
		if err != nil {
			bail("%s.", err)
		}

		err = clipboard.WriteAll(item.Content)
		if err != nil {
			bail("%s.", err)
		}

		fmt.Printf("Item %s copied to the clipboard.\n", paint("green", name))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
