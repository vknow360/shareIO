package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vknow360/shareIO/utils"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func scheduleDelete(file string) {
	time.AfterFunc(time.Minute*5, func() {
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
}
