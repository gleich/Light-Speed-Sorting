package files

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/djherbis/times.v1"
)

// MoveFiles ... Create and move the files
func MoveFiles(files []string) {
	for _, file := range files {
		t, err := times.Stat(file)
		if err != nil {
			logrus.Error(err)
		}
		var fileTime time.Time
		if t.HasBirthTime() {
			fileTime = t.BirthTime()
		} else {
			fileTime = t.ModTime()
		}
		fmt.Println(fileTime)
	}
}
