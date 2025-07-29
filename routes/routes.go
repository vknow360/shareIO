package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vknow360/shareIO/handlers"
	"github.com/vknow360/shareIO/utils"
	"net/http"
	"path/filepath"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("static/*.html")
	r.Static("/static/", "static")

	//TODO: register routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/upload", handlers.HandleUpload)
	r.Static("/uploads/qr", filepath.Join(utils.GetUploadDir(), "qr"))

	r.GET("/files", handlers.GetFilesList)
	r.DELETE("/files/:filename", handlers.DeleteFile)
	r.DELETE("/files", handlers.DeleteAllFiles)

	r.GET("/download/:filename", handlers.DownloadFile)
	return r
}
