package utils

import (
	"os"
	"path/filepath"
)

func GetUploadDir() string {
	return filepath.Join(os.TempDir(), "shareIO_uploads")
}
