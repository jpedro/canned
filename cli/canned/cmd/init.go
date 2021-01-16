package cmd

import (
	"fmt"
	"os"

	"github.com/jpedro/canned"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new can file",
	Run: func(cmd *cobra.Command, args []string) {
		ensurePassword()

		if _, err := os.Stat(canFile); err == nil {
			if err != nil {
				panic(err)
			}
		}

		_, err := canned.InitCan(canFile, canPassword)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Can initialized in file %s.\n", paint("green", canFile))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("force", "f", false, "Force the creation if the files exists.")
}
