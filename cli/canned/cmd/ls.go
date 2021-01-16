package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/jpedro/canned"
	"github.com/jpedro/tablelize"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Shows all secrets",
	Run: func(cmd *cobra.Command, args []string) {
		ensureFile()
		ensurePassword()
		can, err := canned.OpenCan(canFile, canPassword)
		if err != nil {
			panic(err)
		}

		list(can)
	},
}

func list(can *canned.Can) {
	var data [][]string

	data = append(data, []string{"NAME", "LENGTH", "CREATED", "UPDATED", "TAGS"})
	zero := time.Time{}

	for key, item := range can.Items {
		updated := ""
		if item.Metadata.UpdatedAt != zero {
			updated = item.Metadata.UpdatedAt.Format("2006-01-01")
		}

		data = append(data, []string{
			key,
			fmt.Sprintf("%v", len(item.Content)),
			item.Metadata.CreatedAt.Format("2006-01-01"),
			updated,
			strings.Join(item.Tags, " ")})
	}
	tablelize.Rows(data)
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
