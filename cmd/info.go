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

		infoFile := canFile
		fmt.Println("canFile:", canFile)
		if infoFile == canFiles[0] {
			infoFile = paint("pale", infoFile)
		} else if infoFile == "" {
			infoFile = paint("pale", "(not set) using "+canFiles[0])
		} else {
			infoFile = paint("yellow", infoFile)
		}

		infoPassword := paint("pale", "(not set)")
		if canPassword != "" {
			mod := len(canPassword) / 10
			if mod < 1 {
				mod = 1
			}
			strength := strings.Repeat("*", mod)
			infoPassword = paint("yellow", strength)
		}

		infoVerbose := ""
		if canVerbose {
			infoVerbose = paint("yellow", "true")
		} else {
			infoVerbose = paint("pale", "false")
		}

		fmt.Printf("INFO\n")
		fmt.Printf("    Default files:  %v\n", paint("pale", canFiles))
		fmt.Printf("    Version:        %s\n", canVersion)
		fmt.Println("")
		fmt.Printf("ENVIRONMENT\n")
		fmt.Printf("    File:           %v\n", infoFile)
		fmt.Printf("    Password:       %v\n", infoPassword)
		fmt.Printf("    Verbose:        %v\n", infoVerbose)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
