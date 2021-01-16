package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/jpedro/canned"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an item",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ensureFile()
		ensurePassword()
		name := args[0]
		can, err := canned.OpenCan(canFile, canPassword)
		if err != nil {
			panic(err)
		}

		item, err := can.GetItem(name)
		if err != nil {
			panic(err)
		}

		err = clipboard.WriteAll(item.Content)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Item %s copied to the clipboard.\n", paint("green", name))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
