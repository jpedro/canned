package cmd

import (
    "fmt"
    "time"
    "strings"

    "github.com/jpedro/can"
    "github.com/spf13/cobra"
    "github.com/jpedro/tablelize"
)

var lsCmd = &cobra.Command{
    Use:   "ls",
    Short: "Shows all secrets",
    Run: func(cmd *cobra.Command, args []string) {
        store, err := can.OpenStore(CAN_FILE)
        if err != nil {
            panic(err)
        }

        list(store)
    },
}

func list(store *can.Store) {
    var data [][]string

    data = append(data, []string{"NAME", "LENGTH", "CREATED", "UPDATED", "TAGS"})
    zero := time.Time{}

    for key, item := range store.Items {
        updated := ""
        if item.Metadata.UpdatedAt != zero {
            updated = item.Metadata.UpdatedAt.Format("2006-01-01")
        }

        data = append(data, []string{
            key,
            fmt.Sprintf("%v", len(item.Value)),
            item.Metadata.CreatedAt.Format("2006-01-01"),
            updated,
            strings.Join(item.Tags, " ")})
    }
    tablelize.Rows(data)
}

func init() {
    rootCmd.AddCommand(lsCmd)
}
