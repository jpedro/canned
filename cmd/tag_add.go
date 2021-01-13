package cmd

import (
    "fmt"
    "strings"

    "github.com/spf13/cobra"
)

var tagAddTag string
var tagAddName string

var tagAddCmd = &cobra.Command{
    Use:   "add",
    Short: "Adds TAG to NAME",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("==> TODO: tag add TAG NAME")
        fmt.Println("==> ARGS: " + strings.Join(args, " "))
    },
}

func init() {
    tagCmd.AddCommand(tagAddCmd)
}
