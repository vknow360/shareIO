package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/vknow360/shareIO/utils"
)

func isDir(path string) bool {
	fInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fInfo.IsDir()
}

func filesList() []string {
	var files []string = make([]string, 0)
	dst := utils.GetUploadDir()
	f, err := os.Open(dst)
	if err != nil {
		return files
	}
	fileInfo, err := f.Readdirnames(0)
	if err != nil {
		return files
	}
	for _, file := range fileInfo {
		if isDir(filepath.Join(dst, file)) {
			continue
		} else {
			files = append(files, file)
		}
	}
	return files
}

func GetFilesList(c *gin.Context) {
	files := filesList()
	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

func DeleteFile(c *gin.Context) {
	file := c.Param("filename")
	//decode uri component
	file, err := url.PathUnescape(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file name"})
		return
	}
	dst := utils.GetUploadDir()
	filePath := filepath.Join(dst, file)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found: " + err.Error()})
		return
	}

	if err := os.Remove(filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete file"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteAllFiles(c *gin.Context) {
	dst := utils.GetUploadDir()
	files, err := os.ReadDir(dst)

	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusOK, gin.H{"message": "No files to delete"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read directory: " + err.Error()})
		return
	}

	var errors []string
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(dst, file.Name())
			if err := os.Remove(filePath); err != nil {
				errors = append(errors, fmt.Sprintf("Failed to delete %s: %v", file.Name(), err))
			}
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusPartialContent, gin.H{
			"message": "Some files could not be deleted",
			"errors":  errors,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All files deleted successfully"})
}
