package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/jpedro/color"
)

var (
	USER_HOME string
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

func expandHome(path string) string {
	if USER_HOME == "" {
		current, _ := user.Current()
		USER_HOME = current.HomeDir
	}

	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(USER_HOME, path[2:])
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
