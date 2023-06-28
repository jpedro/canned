package cmd

import (
	"fmt"
	"strconv"
	"time"
	"rand"

	"github.com/spf13/cobra"
)

const (
	LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generates a random text",
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		length := strconv.ParseInt(args[0], 10, 0)
		letters := len(LETTERS)
		buffer := make([]rune, length)
		for i := range length {
			buffer[i] = LETTERS[rand.Intn(letters)]
		}
		fmt.Println(string(buffer))
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
