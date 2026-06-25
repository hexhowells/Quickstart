package cmd

import (
	"fmt"
	"os"

	"quickstart/internal/pages"
	"quickstart/internal/ui"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "quickstart [tool]",
  Short: "Quickstart is a minimal CLI tool for rapid documentation",
  Long: `A fast, local documentation viewer and manager like man pages but shorter.`,
  Args:  cobra.MaximumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) == 0 {
		cmd.Help()
		return
	}
	
	toolName := args[0]
	
	filePath, err := pages.GetPath(toolName)
	if err != nil {
		fmt.Printf("[FATAL] - Error resolving path: %v\n", err)
		return
	}

	pageExists := pages.PageExists(filePath)

	if !pageExists {
		fmt.Printf("No quickstart page exists for '%s'\n", toolName)
		fmt.Printf("Run 'quickstart new %s' to create one.\n", toolName)
		return
	}

	if err := ui.RenderPage(filePath); err != nil {
		fmt.Printf("[FATAL] - Error rendering page: %v\n", err)
	}
  },
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {

}