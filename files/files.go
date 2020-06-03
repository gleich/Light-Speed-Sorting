package files

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// GetFiles ... Get all files recursively
func GetFiles() []string {
	files := []string{}
	err := filepath.Walk(".", func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			logrus.Fatal(err)
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		logrus.Fatal(err)
	}
	return files
}
