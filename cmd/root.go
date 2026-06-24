package cmd

import (
	"fmt"
	"os"
	
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
	fmt.Printf("🔍 Searching and displaying quickstart for: %s\n", toolName)
	// TODO: call internal/ui.RenderPage(toolName)
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