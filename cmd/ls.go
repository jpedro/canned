package cmd

import (
	"strings"
	"time"

	"github.com/jpedro/tablelize"
	"github.com/spf13/cobra"

	"github.com/jpedro/canned/lib"
)

type lsOptions struct {
	output string
}

func newLsCommand() *cobra.Command {
	options := lsOptions{}

	cmd := &cobra.Command{
		Use:   "ls",
		Short: "Lists all secrets",
		Run: func(cmd *cobra.Command, args []string) {
			ensureFileExists()
			ensureWeHaveThePassword()

			can, err := canned.OpenCan(canFile, canPassword)
			if err != nil {
				bail("%s.", err)
			}

			listItems(can)
		},
	}

	cmd.Flags().StringVarP(&options.output, "output", "o", "table", "The outout format")

	return cmd
}

func listItems(can *canned.Can) {
	var data [][]string

	data = append(data, []string{"ITEM", "STRENGTH", "CREATED", "UPDATED", "TAGS"})
	zero := time.Time{}

	for key, item := range can.Items {
		created := item.Metadata.CreatedAt.Format("2006-01-01")
		updated := ""
		if item.Metadata.UpdatedAt != zero {
			updated = item.Metadata.UpdatedAt.Format("2006-01-01")
		}

		mod := len(item.Content) / 10
		if mod < 1 {
			mod = 1
		}
		// strength := 1
		data = append(data, []string{
			key,
			strings.Repeat("*", mod),
			created,
			updated,
			strings.Join(item.Tags, " ")})
	}
	tablelize.Rows(data)
}

func init() {
	rootCmd.AddCommand(newLsCommand())
}
