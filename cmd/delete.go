package cmd

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"

	"quickstart/internal/pages"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [tool]",
	Short: "Delete a quickstart page",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toolName := args[0]
		filePath, err := pages.GetPathPath(toolName)

		if err != nil {
			log.Fatalf("[FATAL] - Error resolving the filepath: %w", err)
		}

		pageExists := pages.PageExists(filePath)

		if !pageExists {
			log.Fatalf("Quickstart page for %s does not exist.", toolName)
		}
		
		fmt.Printf("You are about to delete the quickstart page for %s", toolName)
		fmt.Printf("\nIf you are sure you want to delete this page, type '%s': ", toolName)

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := strings.TrimSpace(strings.ToLower(scanner.Text()))

			if input == toolName {
				err := pages.DeletePage(filePath)
				if err != nil {
					log.Fatalf("[FATAL] - Error deleting file: %w", err)
				}
				fmt.Printf("Quickstart page for %s successfully deleted!\n", toolName)
			} else {
				fmt.Printf("Deletion cancelled.\n")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}