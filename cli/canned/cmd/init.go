package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jpedro/canned"
	"github.com/spf13/cobra"
)

type initOptions struct {
	force bool
}

func newInitCmd() *cobra.Command {
	opt := initOptions{}

	cmd := &cobra.Command{
		Use:   "init [OPTIONS]",
		Short: "Initializes a new can file",
		Run: func(cmd *cobra.Command, args []string) {
			if canFile == "" {
				canFile = canFiles[0]
			}
			// fmt.Println(canFile)
			// fmt.Println(opt)

			if _, err := os.Stat(canFile); err == nil {
				if opt.force {
					fmt.Printf("Overridding file file %s.\n", paint("green", canFile))
				} else {
					bail("File %s already exists. Use the '--force' flag to force a recreation.\n",
						paint("green", canFile))
				}
			} else {
				fmt.Printf("Will initialize a new can in file %s.\n",
					paint("green", canFile))
			}

			ensurePassword()
			dirName := filepath.Dir(canFile)
			os.MkdirAll(dirName, 0700)
			_, err := canned.InitCan(canFile, canPassword)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Can initialized in file %s.\n", paint("green", canFile))
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opt.force, "force", "", false, "Force an override if the current files exists")

	return cmd
}

func init() {
	rootCmd.AddCommand(newInitCmd())
}
