package files

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
)

// GetFiles ... Get all files recursively
func GetFiles() []string {
	var location string = "."

	_, err := os.Stat("new-files")
	if !os.IsNotExist(err) {
		location = "./new-files"
	}

	filePaths := []string{}
	err = filepath.Walk(location, func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			statuser.Error("Failed to get all files recursively", err, 1)
		}
		filePaths = append(filePaths, path)
		return nil
	})
	if err != nil {
		statuser.Error("Failed to get all files recursively", err, 1)
	}

	// Removing all . files and all folders
	cleanedFilePaths := []string{}
	for _, filePath := range filePaths {
		fileOrDir, err := os.Stat(filePath)
		if err != nil {
			statuser.Error("Failed to get status of file: "+filePath, err, 1)
		}
		if !strings.HasPrefix(filePath, ".") && !fileOrDir.IsDir() {
			cleanedFilePaths = append(cleanedFilePaths, filePath)
		}
	}

	return cleanedFilePaths
}
