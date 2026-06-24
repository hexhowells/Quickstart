package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [tool]",
	Short: "Create a new quickstart page",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toolName := args[0]
		filePath, err := pages.GetPathPath(toolName)

		if err != nil {
			log.Fatalf("[FATAL] - Error resolving the filepath: %w", err)
		}

		pages.OpenInEditor(filePath)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}