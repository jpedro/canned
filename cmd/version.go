package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Shows the version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(canVersion)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
