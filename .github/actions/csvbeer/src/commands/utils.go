package commands

import (
	"fmt"
	"os"
	"path/filepath"
)


func loadFile(filePath string) (*os.File, error) {

	filePath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("file not found %v: %w", filePath, err)
	}

	_, err = os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("file not found %v: %w", filePath, err)
	}
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found %v", filePath)
	}
	ext := filepath.Ext(filePath)
	if ext != CsvExt {
		return nil, fmt.Errorf("%v not a %v file", filePath, CsvExt)
	}
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open %v: %w", filePath, err)
	}
	return file, nil
}

