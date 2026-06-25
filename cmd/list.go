package cmd

import (
	"fmt"

	"quickstart/internal/pages"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available quickstart pages",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		pages := pages.GetPages()
		
		for _, pageName := range pages {
			fmt.Println(pageName)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}