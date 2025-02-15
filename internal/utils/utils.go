package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ValidateFileFormat(filePath string, allowedFormats []string) error {
	ext := strings.ToLower(filepath.Ext(filePath))
	for _, format := range allowedFormats {
		if ext == format {
			return nil
		}
	}
	return fmt.Errorf("format file tidak didukung. Gunakan file dengan format: %v", allowedFormats)
}
