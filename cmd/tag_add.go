package cmd

import (
    "fmt"
    "../lib"

    "github.com/spf13/cobra"
)

var tagAddCmd = &cobra.Command{
    Use:   "add",
    Short: "Adds TAG to NAME",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
		tag  := args[0]
		name := args[1]

		can, err := lib.Open(CAN_FILE)
        if err != nil {
            panic(err)
        }

		ok := can.AddTag(name, tag)
        if ok != true {
            bail("==> Tag %s was not added to %s.\n",
                paint("green", tag),
                paint("green", name))
		}

		can.Save()
		fmt.Printf("==> Tag %s was added to %s\n",
			paint("green", tag),
			paint("green", name))
    },
}

func init() {
    tagCmd.AddCommand(tagAddCmd)
}
