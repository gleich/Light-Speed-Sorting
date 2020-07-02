package files

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Matt-Gleich/statuser/v2"
	"gopkg.in/djherbis/times.v1"
)

// MoveFiles ... Create and move the files
func MoveFiles(files []string) {
	os.Mkdir("TEMPFILESTORAGE", os.ModePerm)
	for _, file := range files {
		fileName := strings.Split(file, "/")[len(strings.Split(file, "/"))-1]
		tempStorage := "./TEMPFILESTORAGE/" + fileName
		err := os.Rename(file, tempStorage)
		if err != nil {
			statuser.Error("Failed move file", err, 1)
		}
		t, err := times.Stat(tempStorage)
		if err != nil {
			statuser.Error("Failed to get status for "+tempStorage, err, 1)
		}
		var fileTime time.Time
		if t.HasBirthTime() {
			fileTime = t.BirthTime()
		} else {
			fileTime = t.ModTime()
		}
		day := fileTime.Day()
		var suffix string
		if day > 3 && day < 21 {
			suffix = "th"
		}
		switch day % 10 {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		default:
			suffix = "th"
		}
		directory := fmt.Sprintf("./%v/%v/%v-%v%v/", fileTime.Year(), fileTime.Month(), fileTime.Weekday(), fileTime.Day(), suffix)
		if _, err := os.Stat(directory); os.IsNotExist(err) {
			os.MkdirAll(directory, os.ModePerm)
		}
		os.Rename("./TEMPFILESTORAGE/"+fileName, directory+fileName)
		os.Remove("TEMPFILESTORAGE")
	}
}
