package utils

import (
	"fmt"
	"log"
	"os"
)

func WriteFormattedFile(data []byte, filename, fileType string) error {

	if fileType == "" || filename == "" {
		return fmt.Errorf("filename and file_type cannot be empty")
	}

	fullPath := fmt.Sprintf("%s.%s", filename, fileType)

	err := os.WriteFile(fullPath, data, 0666)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fullPath, err)
	}

	log.Printf("File %s successfully created", fullPath)
	return nil
}
