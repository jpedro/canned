package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
    Use:   "status",
    Short: "Shows the environment status",
    Run: func(cmd *cobra.Command, args []string) {
        file := CAN_FILE
        password := CAN_PASSWORD
        verbose := CAN_VERBOSE
        verbosity := ""

        if file == CAN_FILES[0] {
            file = paint("pale", file)
        } else {
            file = paint("yellow", file)
        }
        if password == "" {
            password = paint("pale", "(not set)")
        } else {
            password = paint("yellow", "****")
        }
        if verbose == false {
            verbosity = paint("pale", "false")
        } else {
            verbosity = paint("yellow", "true")
        }

        fmt.Printf("ENVIRONMENT\n")
        fmt.Printf("  File:           %v\n", file)
        fmt.Printf("  Password:       %v\n", password)
        fmt.Printf("  Verbosity:      %v\n", verbosity)
        fmt.Printf("  Directories:    %v\n", paint("pale", CAN_DIRS))
        fmt.Printf("  Files:          %v\n", paint("pale", CAN_FILES))
    },
}

func init() {
    rootCmd.AddCommand(statusCmd)
}
