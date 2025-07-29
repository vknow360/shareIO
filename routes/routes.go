package routes

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/vknow360/shareIO/handlers"
	"github.com/vknow360/shareIO/utils"
)

func RegisterRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.LoadHTMLGlob("static/*.html")
	r.Static("/static/", "static")

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
