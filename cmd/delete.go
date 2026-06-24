package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [tool]",
	Short: "Delete a quickstart page",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toolName := args[0]
		fmt.Printf("Deleting quickstart page for: %s\n", toolName)
		// TODO: call internal/pages.DeletePage(toolName)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}