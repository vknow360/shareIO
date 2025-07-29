package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/vknow360/shareIO/routes"
	"github.com/vknow360/shareIO/utils"
)

const (
	version   = "1.0.0"
	buildDate = "2025/07/29"
)

//go:embed static/*.html
var staticTemplates embed.FS

//go:embed static/*
var staticFiles embed.FS

func main() {

	port := flag.String("port", "8000", "Server port")
	bind := flag.String("bind", "0.0.0.0", "Bind address (0.0.0.0 for all interfaces, 127.0.0.1 for localhost only)")
	deleteAfter := flag.Duration("delete-after", 5*time.Minute, "Auto-delete files after duration (e.g., 5m, 1h, 30s)")
	maxFileSize := flag.Int64("max-file-size", 100, "Maximum file size in MB")
	uploadDir := flag.String("upload-dir", "", "Upload directory (default: system temp directory)")
	showHelp := flag.Bool("help", false, "Show help information")
	showVersion := flag.Bool("version", false, "Show version information")

	flag.Parse()

	if *showHelp {
		fmt.Println("ShareIO - Local Network File Sharing")
		fmt.Println("Usage:")
		flag.PrintDefaults()
		fmt.Println("\nExamples:")
		fmt.Println("  shareio --port 9000")
		fmt.Println("  shareio --bind 127.0.0.1 --delete-after 10m")
		fmt.Println("  shareio --max-file-size 500")
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("ShareIO v%s\n", version)
		fmt.Printf("Build date: %s\n", buildDate)
		os.Exit(0)
	}

	if envPort := os.Getenv("SHAREIO_PORT"); envPort != "" {
		*port = envPort
	}
	if envUploadDir := os.Getenv("SHAREIO_UPLOAD_DIR"); envUploadDir != "" {
		*uploadDir = envUploadDir
	}
	if envMaxSize := os.Getenv("SHAREIO_MAX_FILE_SIZE"); envMaxSize != "" {
		if size, err := strconv.ParseInt(envMaxSize, 10, 64); err == nil {
			*maxFileSize = size
		}
	}

	setGlobalConfig(*deleteAfter, *maxFileSize, *uploadDir)
	ip := utils.GetLocalIP()

	fmt.Printf("ShareIO v%s starting...\n", version)
	fmt.Printf("Upload directory: %s\n", utils.GetUploadDir())
	fmt.Printf("Files auto-delete after: %s\n", *deleteAfter)
	fmt.Printf("Maximum file size: %d MB\n", *maxFileSize)
	fmt.Printf("Server bind address: %s\n", *bind)
	fmt.Printf("Server port: %s\n", *port)
	fmt.Println("\nAccess URLs:")
	fmt.Printf("  Local:   http://127.0.0.1:%s\n", *port)
	if *bind == "0.0.0.0" {
		fmt.Printf("  Network: http://%s:%s\n", ip, *port)
	}
	fmt.Println("\nPress Ctrl+C to stop the server")

	r := routes.RegisterRoutes(staticTemplates, staticFiles)

	err := r.Run(fmt.Sprintf("%s:%s", *bind, *port))

	if err != nil {
		os.Exit(-1)
	}
}

func setGlobalConfig(deleteAfter time.Duration, maxFileSize int64, uploadDir string) {
	os.Setenv("SHAREIO_DELETE_AFTER", deleteAfter.String())
	os.Setenv("SHAREIO_MAX_FILE_SIZE_BYTES", strconv.FormatInt(maxFileSize*1024*1024, 10))
	if uploadDir != "" {
		os.Setenv("SHAREIO_UPLOAD_DIR", uploadDir)
	}
}
