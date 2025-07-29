package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vknow360/shareIO/utils"
)

func getDeleteAfterDuration() time.Duration {
	if duration := os.Getenv("SHAREIO_DELETE_AFTER"); duration != "" {
		if d, err := time.ParseDuration(duration); err == nil {
			return d
		}
	}
	return 5 * time.Minute
}

func getMaxFileSize() int64 {
	if size := os.Getenv("SHAREIO_MAX_FILE_SIZE_BYTES"); size != "" {
		if s, err := strconv.ParseInt(size, 10, 64); err == nil {
			return s
		}
	}
	return 100 * 1024 * 1024
}

func scheduleDelete(file string) {
	time.AfterFunc(getDeleteAfterDuration(), func() {
		_ = os.Remove(file)
	})
}

func HandleUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("File not received %s", err),
		})
		return
	}

	if file.Size > getMaxFileSize() {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{
			"message": fmt.Sprintf("File too large. Maximum size: %d MB", getMaxFileSize()/(1024*1024)),
		})
		return
	}

	dst := utils.GetUploadDir()
	err = os.MkdirAll(dst, os.ModePerm)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error creating folder: %v", err),
		})
		return
	}

	if err := c.SaveUploadedFile(file, filepath.Join(dst, file.Filename)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error saving file: %v", err),
		})
		return
	}
	go scheduleDelete(filepath.Join(dst, file.Filename))
	c.JSON(http.StatusOK, gin.H{
		"message":      "File uploaded successfully",
		"filename":     file.Filename,
		"size":         file.Size,
		"delete_after": getDeleteAfterDuration().String(),
	})
}
