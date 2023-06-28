package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/jpedro/color"
	"golang.org/x/term"
)

var (
	userHome string
)

func env(name string, fallback string) string {
	value := os.Getenv(name)
	if value == "" {
		return fallback
	}
	return value
}

func paint(name string, text interface{}) string {
	return color.Paint(name, fmt.Sprintf("%s", text))
}

func askPassword(prompt string) string {
	// password := ""
	// fmt.Print("EX: " + prompt)
	// fmt.Print("\033[8m") // Hide input
	// fmt.Scan(&password)
	// fmt.Print("\033[28m")

	// return password
	fmt.Print(prompt)
	bytes, err := term.ReadPassword(syscall.Stdin)
	fmt.Println()
	if err != nil {
		return ""
	}

	password := string(bytes)
	return password
}

func expandHome(path string) string {
	if userHome == "" {
		current, _ := user.Current()
		userHome = current.HomeDir
	}

	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(userHome, path[2:])
	}

	return path
}

func getRandomPassword(length int) string {
	randomChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	randomSymbols = []rune("-")
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

	return string(buffer)
}

func warn(message string, args ...any) {
	fmt.Fprintf(os.Stderr, "\033[33;1mWarning: "+message+"\033[0m", args...)
	if message[len(message)-1:] != "\n" {
		fmt.Println("")
	}
}

func bail(message string, args ...any) {
	fmt.Printf("\033[31;1mError: "+message+"\033[0m", args...)
	if message[len(message)-1:] != "\n" {
		fmt.Println("")
	}
	os.Exit(1)
}
