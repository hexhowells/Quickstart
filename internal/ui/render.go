package ui

import (
	"fmt"
	"os"
	"os/exec"
)


func RenderPage(filePath string) error {
	file, err := os.Open(filePath)
	
	if err != nil {
		return fmt.Errorf("failed to read page: %w", err)
	}
	defer file.Close()

	cmd := exec.Command("less", "-X")

	cmd.Stdin = file

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pager execution failed: %w", err)
	}

	return nil
}