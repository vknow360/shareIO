package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vknow360/shareIO/utils"
	"net/http"
	"os"
	"path/filepath"
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
	err := os.RemoveAll(dst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete files: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
