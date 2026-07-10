package pages

import (
	"os"
	"fmt"
	"log"
	"strings"
	"slices"
	"path/filepath"

	"github.com/adrg/xdg"
)


func GetPath(toolName string) (string, error) {
	baseDir := filepath.Join(xdg.DataHome, "quickstart")

	err := os.MkdirAll(baseDir, 0755)  // mkdir -p
	if err != nil {
		return "", err
	}

	return filepath.Join(baseDir, toolName+".md"), nil
}


func PageExists(filePath string) bool {
	_, err := os.Stat(filePath)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}


func GetPages() []string {
	baseDir := filepath.Join(xdg.DataHome, "quickstart")
	pages, err := os.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}

	availablePages := make([]string, len(pages))
	for i, p := range pages {
		toolName := strings.TrimSuffix(p.Name(), ".md")
		availablePages[i] = toolName
	}

	slices.Sort(availablePages)

	return availablePages
}


func DeletePage(filePath string) error {
	err := os.Remove(filePath)

	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}


func RenamePage(filePath string, newFilePath string) error {
	err := os.Rename(filePath, newFilePath)

	if err != nil {
		return fmt.Errorf("failed to rename file: %w", err)
	}

	return nil
}


func CreatePage(filePath string, toolName string) error {
	template := fmt.Sprintf(`==================================================
%s Quickstart Guide
==================================================

Description:
  [Briefly describe what %s does]

Commands:
  command: %s param1 param2

==================================================`, toolName, toolName, toolName)

	err := os.WriteFile(filePath, []byte(template), 0644)

	if err != nil {
		return fmt.Errorf("[FATAL] - Error creating new page: %w\n", err)
	}

	return nil
}