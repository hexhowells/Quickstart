package pages

import (
	"fmt"
	"os"
	"os/exec"
)


func OpenInEditor(filePath string) error {
	editor := os.Getenv("VISUAL")
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}

	if editor == "" {
		editor = "vi"
	}

	cmd := exec.Command(editor, filePath)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run editor %s: %w", editor, err)
	}

	return nil
}