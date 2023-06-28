package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	canPassword string
	canFile     string
	canVerbose  bool

	canFiles = []string{
		expandHome("~/.config/canned/default.can"),
		"/etc/canned/default.can",
	}

	rootCmd = &cobra.Command{
		Use:   "canned",
		Short: "Canned stores encrypted goodies",
		// Run: func(cmd *cobra.Command, args []string) {
		// 	usage(nil, []string{})
		// },
	}
)

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	return nil
}

func usage(cmd *cobra.Command, text []string) {
	fmt.Println(`USAGE
    canned init                 # Initializes a new can file
    canned ls                   # Lists all items
    canned set NAME [VALUE]     # Stores an item
    canned get NAME             # Copies the item's content to the clipboard
    canned rm NAME              # Removes an item
    canned tag ls               # Shows all tags
    canned tag add TAG NAME     # Adds the tag TAG to item NAME
    canned tag rm TAG NAME      # Removes the tag TAG from item NAME
    canned random [LENGTH]      # Generates a new random value
    canned env                  # Shows the environment status
    canned version              # Shows the version
    canned help                 # Shows this help

GLOBAL OPTIONS
    -f, --file FILE             # Use a custom file
    -v, --verbose               # Shows verbose output

ENVIRONMENT VARIABLES
    CANNED_FILE                 # Use this file instead of the default
    CANNED_PASSWORD             # Use this password (avoids the password prompt)
    CANNED_VERBOSE              # Turns verbosity on
    `)

}

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd.SetHelpFunc(usage)
	rootCmd.PersistentFlags().StringVarP(&canFile, "file", "f", "", "Custom can file path")
	rootCmd.PersistentFlags().BoolVarP(&canVerbose, "verbose", "v", false, "Show verbose output")
	// rootCmd.PersistentFlags().BoolVarP(&canOverwrite, "overwrite", "", false, "Init can overwrite the same file")
	log.Printf("INIT %s\n", canFile)
	log.Printf("INIT %t\n", canVerbose)
}

func initConfig() {
	canPassword = env("CANNED_PASSWORD", "")
	canFile = env("CANNED_FILE", "")
	canVerbose, _ = strconv.ParseBool(env("CANNED_VERBOSE", "false"))
	// canOverwrite, _ = strconv.ParseBool(env("CANNED_OVERWRITE", "false"))
	log.Printf("INIT_CONFIG %s\n", canFile)
	log.Printf("INIT_CONFIG %t\n", canVerbose)
}

func ensureFile() {
	if canFile != "" {
		info, err := os.Stat(canFile)
		if os.IsNotExist(err) {
			bail("Error: The file %s does not exist.", paint("green", canFile))
		} else if info.IsDir() {
			bail("Error: The file %s is a directory.", paint("green", canFile))
		} else {
			return
		}
	}

	canFile = findNearestFile()
	if canFile == "" {
		bail("Error: Couldn't find a default can file. Use the '--file FILE' flag or set the 'CANNED_FILE' env var to point to a valid can file.")
	}
}

func findNearestFile() string {
	for _, file := range canFiles {
		info, err := os.Stat(file)
		if os.IsNotExist(err) {
			// continue
		} else if info.IsDir() {
			// continue
		} else {
			return file
		}
	}

	return ""
}

func ensurePassword() {
	if canPassword != "" {
		return
	}

	password := askPassword("Enter the can password: ")
	// fmt.Printf("Enter the password: ")
	// reader := bufio.NewReader(os.Stdin)
	// pass, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)
	if password == "" {
		fmt.Println("Error: Password can't be empty.")
		os.Exit(1)
	}

	canPassword = password
}
