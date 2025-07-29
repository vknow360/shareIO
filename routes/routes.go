package routes

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vknow360/shareIO/handlers"
)

func RegisterRoutes(staticTemplates embed.FS, staticFiles embed.FS) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	tmplFS, err := fs.Sub(staticTemplates, "static")
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(template.Must(template.ParseFS(tmplFS, "*.html")))

	staticAssets, err := fs.Sub(staticFiles, "static")
	if err != nil {
		panic(err)
	}
	r.StaticFS("/static", http.FS(staticAssets))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/upload", handlers.HandleUpload)

	r.GET("/files", handlers.GetFilesList)
	r.DELETE("/files/:filename", handlers.DeleteFile)
	r.DELETE("/files", handlers.DeleteAllFiles)

	r.GET("/download/:filename", handlers.DownloadFile)
	return r
}
