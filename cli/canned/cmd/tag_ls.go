package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var tagLsCmd = &cobra.Command{
    Use:   "ls",
    Short: "Lists all used tags",
    Run: func(cmd *cobra.Command, args []string) {
		ensureFile()
		ensurePassword()
        fmt.Println("==> TODO: tag ls")
    },
}

func init() {
    tagCmd.AddCommand(tagLsCmd)
}
