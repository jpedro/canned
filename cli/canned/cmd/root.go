package cmd

import (
    "fmt"
    "os"
    "strconv"
    "bufio"
    "strings"

    "github.com/spf13/cobra"
)

var CAN_VERBOSE     bool
var CAN_PASSWORD    string
var CAN_FILE        string

var CAN_DIRS = []string{
    expand("~/.config/can"),
    "/etc/can",
}
var CAN_FILES = []string{
    expand("~/.config/can/main2.can"),
    "/etc/can/main2.can",
}

var rootCmd = &cobra.Command{
    Use:   "can",
    Short: "Can stores encrypted goodies",
    Run: func(cmd *cobra.Command, args []string) {
        usage(nil, []string{})
    },
}

func Execute() error {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }

    return nil
}

func usage(cmd *cobra.Command, text []string) {
    fmt.Println(`USAGE
    can init                # Initializes a new can file
    can ls                  # Lists all secrets
    can set NAME [VALUE]    # Lists all secrets
    can get NAME            # Copies the secret value to the clipboard
    can rm NAME             # Removes a secret
    can tag ls              # Shows all tags
    can tag add TAG NAME    # Adds the tag TAG to secret NAME
    can tag rm TAG NAME     # Removes the tag TAG from secret NAME
    can random [LENGTH]     # Generates a new random value
    can status              # Shows the environment status
    can version             # Shows the version
    can help                # Shows this help

GLOBAL OPTIONS
    -f, --file FILE          # Use a custom file
    -v, --verbose            # Shows verbose output

ENVIRONMENT VARIABLES
    CAN_FILE                 # Use this file instead of the default
    CAN_PASSWORD             # Use this password (avoids the password prompt)
    CAN_VERBOSE              # Turns verbosity on
    CAN_AUTO_INIT            # Initializes the can file if it's not ready
    `)

}

func init() {
    // rootCmd.SetHelpFunc(usage)
    rootCmd.PersistentFlags().BoolVarP(&CAN_VERBOSE, "verbose", "v", false, "Show verbose output")
    rootCmd.PersistentFlags().StringVarP(&CAN_FILE, "file", "f", "", "Can file path")

    CAN_FILE        = env("CAN_FILE", CAN_FILES[0])
    CAN_PASSWORD    = env("CAN_PASSWORD", "")
    CAN_VERBOSE, _  = strconv.ParseBool(env("CAN_VERBOSE", "false"))
}

func ensurePassword() {
    if CAN_PASSWORD != "" {
        return
    }

    fmt.Printf("Enter the password: ")
    reader := bufio.NewReader(os.Stdin)
    pass, _ := reader.ReadString('\n')
    pass = strings.TrimSpace(pass)
    if pass == "" {
        fmt.Println("Error: Password can't be empty.")
        os.Exit(1)
    }

    CAN_PASSWORD = pass
}
