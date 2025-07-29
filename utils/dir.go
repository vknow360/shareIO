package utils

import (
	"os"
	"path/filepath"
)

func GetUploadDir() string {
	if customDir := os.Getenv("SHAREIO_UPLOAD_DIR"); customDir != "" {
		return customDir
	}
	return filepath.Join(os.TempDir(), "shareIO_uploads")
}
