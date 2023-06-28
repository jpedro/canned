package cmd

import (
	"fmt"
	"strings"

	"github.com/jpedro/tablelize"
	"github.com/spf13/cobra"

	"github.com/jpedro/canned/lib"
)

type tagStats struct {
	Count int
	Items []string
}

// var tagsStats map[string]*tagStats

var tagLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all used tags",
	Run: func(cmd *cobra.Command, args []string) {
		ensureFileExists()
		ensureWeHaveThePassword()

		can, err := canned.OpenCan(canFile, canPassword)
		if err != nil {
			bail("%s.", err)
		}

		tagsStats := make(map[string]*tagStats)
		for name, item := range can.Items {
			// log.Printf("name %s\n", name)
			for _, tag := range item.Tags {
				// log.Printf("tag %s\n", tag)
				stats, ok := tagsStats[tag]
				if !ok {
					stats = &tagStats{0, []string{}}
				}
				// log.Printf("stats %v\n", stats)
				stats.Count = stats.Count + 1
				stats.Items = append(stats.Items, name)
				tagsStats[tag] = stats
				// log.Printf("tagsStats %v\n", tagsStats)
			}
		}
		// log.Printf("tagsStats %v\n", tagsStats)
		var data [][]string
		data = append(data, []string{"COUNT", "TAG", "ITEMS"})

		for name, stats := range tagsStats {
			// log.Printf("stats %v\n", stats)
			data = append(data, []string{
				fmt.Sprintf("%d", stats.Count),
				name,
				strings.Join(stats.Items, " ")},
			)
		}
		tablelize.Rows(data)
	},
}

func init() {
	tagCmd.AddCommand(tagLsCmd)
}
