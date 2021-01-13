package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Shows all secrets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Secret list: %v\n", CAN_VERBOSE)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
