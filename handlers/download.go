package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vknow360/shareIO/utils"
)

func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"filename": "Invalid filename",
		})
		return
	}
	filePath := filepath.Join(utils.GetUploadDir(), filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"filename": filename,
		})
		return
	}

	c.FileAttachment(filePath, filename)
}
