package cmd

import (
    "fmt"

    "github.com/jpedro/can"
    "github.com/spf13/cobra"
)

var tagRmCmd = &cobra.Command{
    Use:   "rm",
    Short: "Removes TAG from NAME",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        tag  := args[0]
        name := args[1]

        store, err := can.OpenStore(CAN_FILE)
        if err != nil {
            panic(err)
        }

        ok := store.DelTag(name, tag)
        if !ok {
            bail("Error: Tag %s was not removed from %s.\n",
                paint("green", tag),
                paint("green", name))
        }

        err = store.Save()
        if err != nil {
            bail("==> Failed to save file %s%s.\n",
                paint("green", CAN_FILE))
        }

        fmt.Printf("==> Tag %s removed from %s\n",
            paint("green", tag),
            paint("green", name))
    },
}

func init() {
    tagCmd.AddCommand(tagRmCmd)
}
