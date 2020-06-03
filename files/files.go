package files

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// GetFiles ... Get all files recursively
func GetFiles() []string {
	filePaths := []string{}
	err := filepath.Walk(".", func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			logrus.Fatal(err)
		}
		filePaths = append(filePaths, path)
		return nil
	})
	if err != nil {
		logrus.Fatal(err)
	}

	// Removing all . files and all folders
	cleanedFilePaths := []string{}
	for _, filePath := range filePaths {
		fileOrDir, err := os.Stat(filePath)
		if err != nil {
			logrus.Fatal(err)
		}
		if !strings.HasPrefix(filePath, ".") && !fileOrDir.IsDir() {
			cleanedFilePaths = append(cleanedFilePaths, filePath)
		}
	}

	return cleanedFilePaths
}
