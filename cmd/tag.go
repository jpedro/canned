package cmd

import (
    // "fmt"

    "github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
    Use:   "tag",
    Short: "Shows all used tags",
    // Run: func(cmd *cobra.Command, args []string) {
    //     tagLsCmd.Execute()
    // },
}

func init() {
    rootCmd.AddCommand(tagCmd)
}
