package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/jpedro/canned/lib"
)
type setOptions struct {
	force bool
}

func newSetCmd() *cobra.Command {
	options := setOptions{}

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Sets a new item NAME VALUE (use '@random' for a random value)",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			ensureFileExists()
			ensureWeHaveThePassword()

			name := args[0]
			value := args[1]
			if value == "@random" {
				value = getRandomPassword(42)
			}
			// var value string
			// if len(args) < 2 {
			// // if value == "" {
			// 	value = getRandomPassword(42)
			// } else {
			// 	value = args[1]
			// }

			can, err := canned.OpenCan(canFile, canPassword)
			if err != nil {
				bail("%s.", err)
			}

			item, _ := can.GetItem(name)
			if item != nil {
				if options.force {
					warn(
						"Overridding item '%s'.\n",
						name,
					)
				} else {
					bail(
						"Item '%s' already exists. Use '--force' to recreate the item anew.",
						name,
					)
				}
			}

			err = can.SetItem(name, value)
			if err != nil {
				bail("%s.", err)
			}

			err = can.Save()
			if err != nil {
				bail("%s.", err)
			}

			fmt.Printf("Item '%s' stored.\n", paint("green", name))
		},
	}

	cmd.Flags().BoolVarP(&options.force, "force", "", false, "Override the existing item")

	return cmd
}

func init() {
	rootCmd.AddCommand(newSetCmd())
}
