package cmd

import (
	"fmt"
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
	if USER_HOME == "" {
		current, _ := user.Current()
		userHome = current.HomeDir
	}

	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(userHome, path[2:])
	}

	return path
}

func bail(format string, msg ...interface{}) {
	fmt.Printf("Error: "+format, msg...)
	if format[len(format)-1:] != "\n" {
		fmt.Println("")
	}
	os.Exit(1)
}
