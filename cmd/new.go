package cmd

import (
	"log"
	"fmt"
	
	"quickstart/internal/pages"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [tool]",
	Short: "Create a new quickstart page",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toolName := args[0]
		filePath, err := pages.GetPath(toolName)

		if err != nil {
			log.Fatalf("[FATAL] - Error resolving the filepath: %w\n", err)
		}

		pageExists := pages.PageExists(filePath)

		if pageExists {
			fmt.Printf("Quickstart page for '%s' already exists.\n", toolName)
			return
		} else {
			pages.CreatePage(filePath, toolName)
			pages.OpenInEditor(filePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}