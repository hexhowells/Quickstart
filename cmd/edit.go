package cmd

import (
	"log"

	"quickstart/internal/pages"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [tool]",
	Short: "Edit an existing quickstart page",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toolName := args[0]
		filePath, err := pages.GetPathPath(toolName)

		if err != nil {
			log.Fatalf("[FATAL] - Error resolving the filepath: %w\n", err)
		}

		pageExists := pages.PageExists(filePath)

		if pageExists {
			pages.OpenInEditor(filePath)
		} else {
			log.Fatalf("Quickstart page for %s does not exist. Please run 'quickstart new %s' first.", toolName, toolName)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}