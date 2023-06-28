package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Shows information and the environment",
	Run: func(cmd *cobra.Command, args []string) {
		file := canFile
		password := canPassword
		verbose := canVerbose
		verbosity := ""

		if file == canFiles[0] {
			file = paint("pale", file)
		} else if file == "" {
			file = paint("pale", "(not set)")
		} else {
			file = paint("yellow", file)
		}

		if password == "" {
			password = paint("pale", "(not set)")
		} else {
			mod := len(canPassword) / 10
			if mod < 1 {
				mod = 1
			}
			strength := strings.Repeat("*", mod)
			password = paint("yellow", strength)
		}

		if verbose {
			verbosity = paint("yellow", "true")
		} else {
			verbosity = paint("pale", "false")
		}

		fmt.Printf("INFO\n")
		fmt.Printf("    Default files:  %v\n", paint("pale", canFiles))
		fmt.Printf("    Version:        %s\n", canVersion)
		fmt.Println("")
		fmt.Printf("ENVIRONMENT\n")
		fmt.Printf("    File:           %v\n", file)
		fmt.Printf("    Password:       %v\n", password)
		fmt.Printf("    Verbosity:      %v\n", verbosity)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
