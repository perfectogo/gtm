package initialize

import (
	"bytes"
	"errors"
	"os"
)

func CreateFile(projectName string, files map[string]string) error {
	for path, code := range files {

		file, err := os.Create(path)
		if err != nil {
			return errors.New("Error creating file. Error " + err.Error())
		}
		defer file.Close()

		var buffer bytes.Buffer
		{
			if err := Tmp(code).Execute(&buffer, nil); err != nil {
				return errors.New("Error executing template. Error: " + err.Error())
			}

			if _, err := buffer.WriteTo(file); err != nil {
				return errors.New("Error writing to file. Error: " + err.Error())
			}
		}

	}

	return nil
}
