package cmd

import (
    "fmt"
    "../lib"

    "github.com/spf13/cobra"
)

var tagRmCmd = &cobra.Command{
    Use:   "rm",
    Short: "Removes TAG from NAME",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        tag  := args[0]
        name := args[1]

        can, err := lib.Open(CAN_FILE)
        if err != nil {
            panic(err)
        }

        ok := can.DelTag(name, tag)
        if ok != true {
            bail("Error: Tag %s was not removed from %s.\n", tag, name)
                // paint("green", tag),
                // paint("green", name))
        }

        can.Save()
        fmt.Printf("==> Tag %s removed from %s\n", paint("green", tag), paint("green", name))
    },
}

func init() {
    tagCmd.AddCommand(tagRmCmd)
}
