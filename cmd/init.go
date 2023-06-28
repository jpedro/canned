package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jpedro/canned"
	"github.com/spf13/cobra"
)

type initOptions struct {
	overwrite bool
}

func newInitCmd() *cobra.Command {
	options := initOptions{}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initializes a new can file",
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("INIT_RUN %s\n", canFile)
			ensureFile()
			log.Printf("INIT_RUN %s\n", canFile)

			if _, err := os.Stat(canFile); err == nil {
				if options.overwrite {
					fmt.Printf("Overridding file file %s.\n", paint("green", canFile))
				} else {
					bail("File %s already exists. Use '--overwrite' to force a new file.\n",
						paint("green", canFile))
				}
			} else {
				fmt.Printf("Will initialize a new can in file %s.\n",
					paint("green", canFile))
			}

			ensurePassword()
			dirName := filepath.Dir(canFile)
			err := os.MkdirAll(dirName, 0700)
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

	cmd.Flags().BoolVarP(&options.overwrite, "overwrite", "", false, "Init can override an existing file")

	return cmd
}

func init() {
	rootCmd.AddCommand(newInitCmd())
}
