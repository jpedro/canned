package cmd

import (
	"fmt"
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

	l := len(LETTERS)
	b := make([]rune, n)
	for i := range b {
			b[i] = LETTERS[rand.Intn(l)]
	}
	return string(b)
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
