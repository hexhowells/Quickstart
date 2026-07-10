package cmd

import (
	"log"

	"quickstart/internal/pages"

	"github.com/spf13/cobra"
)


var renameCmd = &cobra.Command{
	Use:   "rename [tool] [tool]",
	Short: "Rename a quickstart page",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		toolName := args[0]
		newToolName := args[1]
		filePath, err := pages.GetPath(toolName)

		if err != nil {
			log.Fatalf("[FATAL] - Error resolving the filepath: %w\n", err)
		}

		newFilePath, err := pages.GetPath(newToolName)

		if err != nil {
			log.Fatalf("[FATAL] - Error resolving the filepath: %w\n", err)
		}

		pageExists := pages.PageExists(filePath)
		newPageAlreadyExists := pages.PageExists(newFilePath)

		if newPageAlreadyExists {
			log.Fatalf("Quickstart page with name %s already exists! Please delete the page or use a different name.", newToolName)
		}

		if pageExists {
			pages.RenamePage(filePath, newFilePath)
		} else {
			log.Fatalf("Quickstart page for %s does not exist. Please check the spelling. " +
			"Additionally, you can check existing tools with \"quickstart list\"", 
			toolName)
		}
	},
}


func init() {
	rootCmd.AddCommand(renameCmd)
}