package cmd

import (
	"fmt"
	"math/rand"
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

			totalChars := len(randomChars)
			totalSymbols := len(randomSymbols)
			buffer := make([]rune, length)
			for i := range buffer {
				if i > 0 && (i + 1) % 6 == 0 && i < (length -1) {
					buffer[i] = randomSymbols[rand.Intn(totalSymbols)]
				} else {
					buffer[i] = randomChars[rand.Intn(totalChars)]
				}
			}
			fmt.Println(string(buffer))
		},
	}
)

func init() {
	rootCmd.AddCommand(randomCmd)
}
