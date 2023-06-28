package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/jpedro/canned/lib"
)

type initOptions struct {
	force bool
}

func newInitCmd() *cobra.Command {
	options := initOptions{}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initializes a new can file",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := os.Stat(canFile)
			if err == nil {
				if options.force {
					fmt.Printf(
						"Overridding file file %s.\n",
						paint("green", canFile),
					)
				} else {
					bail(
						"File %s already exists. Use '--force' to recreate the file anew.",
						paint("green", canFile),
					)
				}

			} else {
				fmt.Printf(
					"Will initialize can in a fresh file %s.\n",
					paint("green", canFile),
				)
			}

			ensureWeHaveThePassword()
			dirName := filepath.Dir(canFile)
			err = os.MkdirAll(dirName, 0700)
			if err != nil {
				panic(err)
			}

			_, err = canned.InitCan(canFile, canPassword)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Can initialized in file %s.\n", paint("green", canFile))
		},
	}

	cmd.Flags().BoolVarP(&options.force, "force", "", false, "Override the existing file")

	return cmd
}

func init() {
	rootCmd.AddCommand(newInitCmd())
}
