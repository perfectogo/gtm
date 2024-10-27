package initialize

import (
	"errors"
	"os"
)

func CreateFolder(folders []string) error {
	for _, folder := range folders {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			return errors.New("Error creating folder: " + folder)
		}
	}

	return nil
}
