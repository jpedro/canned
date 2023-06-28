package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	randomChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	randomSymbols = []rune("-")
	randomCmd = &cobra.Command{
		Use:   "random",
		Short: "Generates a random text",
		Run: func(cmd *cobra.Command, args []string) {
			length := 40

			if len(args) > 0 {
				val, err := strconv.Atoi(args[0])
				if err != nil {
					bail("Invalid length '%s'.", args[0])
					return
				}
				length = val
			}

			fmt.Println(getRandomPassword(length))
		},
	}
)

func init() {
	rootCmd.AddCommand(randomCmd)
}
