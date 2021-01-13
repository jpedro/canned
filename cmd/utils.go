package cmd

import (
    // "log"
    "os"
    "os/user"
    "strings"
    "path/filepath"

    "github.com/jpedro/color"
)

var USER_HOME string

func env(name string, fallback string) string {
    value := os.Getenv(name)
    if value == "" {
        return fallback
    }
    return value
}

func paint(name string, text interface{}) string {
    return color.Paint(name, text)
}

func expand(path string) string {
    if USER_HOME == "" {
        usr, _ := user.Current()
        USER_HOME = usr.HomeDir
        // log.Printf("==> Resolving %v.\n", USER_HOME)
    }

    if strings.HasPrefix(path, "~/") {
        path = filepath.Join(USER_HOME, path[2:])
    }

    return path
}
