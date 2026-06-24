package pages

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)


func GetPathPath(toolName string) (string, error) {
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
