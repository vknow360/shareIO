# shareIO - Local Network File Sharing

shareIO is a lightweight, cross-platform file sharing server designed for quick and easy file transfers between devices on the same local network. Perfect for offline environments where you need to share files between computers, phones, tablets, and other devices without internet connectivity.

## Features

- 🚀 **Zero Configuration** - Just run and share
- 🌐 **Web-based Interface** - Works on any device with a browser
- 📱 **Cross-platform** - Windows, macOS, Linux support
- ⚡ **Configurable Auto-cleanup** - Files auto-delete after specified time (default: 5 minutes)
- 🔒 **Local Network Only** - No internet required, keeps files private
- 🗂️ **File Management** - Upload, download, delete individual or all files
- 🎛️ **Flexible Configuration** - Command line options and environment variables
- 🔧 **Size Limits** - Configurable maximum file size (default: 100MB)

## Quick Start

### For End Users

1. **Download** the binary for your operating system
2. **Run** the application:
   - Windows: `shareIO.exe`
   - macOS/Linux: `./shareIO`
3. **Access** the web interface at the displayed URLs
4. **Share** files between devices on your network

### Basic Usage

1. Open your web browser and go to the displayed URL
2. Upload files using the web interface
3. Other devices can access the same URL to view and download files
4. Files automatically delete after the configured time (default: 5 minutes)
5. Use delete buttons to remove files immediately

## Command Line Options

```bash
# Basic usage - starts on port 8000, accessible from network
./shareIO

# Custom port
./shareIO --port 9000

# Localhost only (more secure)
./shareIO --bind 127.0.0.1

# Change auto-delete time (supports: s, m, h)
./shareIO --delete-after 10m
./shareIO --delete-after 1h
./shareIO --delete-after 30s

# Set maximum file size in MB
./shareIO --max-file-size 500

# Custom upload directory
./shareIO --upload-dir /path/to/uploads

# Show help
./shareIO --help

# Show version information
./shareIO --version
```

### Environment Variables

You can also configure shareIO using environment variables:

```bash
# Set port
export shareIO_PORT=9000

# Set upload directory
export shareIO_UPLOAD_DIR=/path/to/uploads

# Set max file size in MB
export shareIO_MAX_FILE_SIZE=200

./shareIO
```

## Configuration Examples

### Home Network File Sharing
```bash
# Default settings - accessible to all network devices
./shareIO
```

### Secure Local Development
```bash
# Localhost only access
./shareIO --bind 127.0.0.1 --port 3000
```

### Conference Room Setup
```bash
# Longer file retention, larger files allowed
./shareIO --delete-after 2h --max-file-size 1000 --port 8080
```

### Custom Storage Location
```bash
# Use specific directory for uploads
./shareIO --upload-dir ./shared_files --delete-after 30m
```

## Network Access

When shareIO starts, it displays:
```
shareIO v1.0.0 starting...
Upload directory: C:\Users\...\Temp\shareIO_uploads
Files auto-delete after: 5m0s
Maximum file size: 100 MB
Server bind address: 0.0.0.0
Server port: 8000

Access URLs:
  Local:   http://127.0.0.1:8000
  Network: http://192.168.1.100:8000

Press Ctrl+C to stop the server
```

Share the **Network URL** with other devices to allow file sharing access.

## Security Considerations

shareIO is designed for **trusted local networks**:

- ✅ **Safe for**: Home networks, office LANs, conference rooms, offline environments
- ⚠️ **Use caution**: Public WiFi, untrusted networks
- 🔒 **Security features**: 
  - Automatic file cleanup after specified time
  - Configurable file size limits
  - Path traversal protection in downloads
  - Local network binding options
  - No permanent file storage by default

## File Operations

### Via Web Interface
- **Upload**: Select and upload files through the browser
- **Download**: Click on files in the list to download
- **Delete**: Remove individual files or all files at once
- **View**: See all available files with their names

### File Lifecycle
1. **Upload** - Files saved to upload directory
2. **Access** - Available for download from any network device
3. **Auto-delete** - Automatically removed after configured time
4. **Manual cleanup** - Can be deleted immediately via web interface

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/` | Web interface homepage |
| `POST` | `/upload` | Upload file endpoint |
| `GET` | `/files` | Get list of available files (JSON) |
| `GET` | `/download/:filename` | Download specific file |
| `DELETE` | `/files/:filename` | Delete specific file |
| `DELETE` | `/files` | Delete all files |
| `GET` | `/static/*` | Static assets (CSS, JS, images) |

### API Response Examples

**File Upload Response:**
```json
{
  "message": "File uploaded successfully",
  "filename": "document.pdf",
  "size": 1048576,
  "delete_after": "5m0s"
}
```

**Files List Response:**
```json
{
  "files": ["document.pdf", "image.jpg", "data.csv"]
}
```

## Building from Source

### Prerequisites
- Go 1.19 or later
- Git

### Build Steps
```bash
git clone https://github.com/vknow360/shareIO.git
cd shareIO
go mod tidy
go build -o shareIO .
```

### Cross-compilation
```bash
# Windows 64-bit
GOOS=windows GOARCH=amd64 go build -o shareIO-windows-amd64.exe .

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o shareIO-darwin-amd64 .

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o shareIO-darwin-arm64 .

# Linux 64-bit
GOOS=linux GOARCH=amd64 go build -o shareIO-linux-amd64 .
```

## Project Structure

```
shareIO/
├── main.go                    # Application entry point & configuration
├── routes/
│   └── routes.go             # HTTP route registration
├── handlers/
│   ├── upload.go             # File upload logic
│   ├── download.go           # File download with security checks
│   └── files.go              # File listing and deletion
├── utils/
│   ├── network.go            # Local IP detection
│   └── dir.go                # Upload directory management
├── static/
│   ├── *.html                # Web interface templates
│   └── assets/               # CSS, JS, images
├── go.mod                    # Go module definition
└── README.md                 # This file
```

## Troubleshooting

### Common Issues

**"Port already in use" error**
```bash
./shareIO --port 8001
```

**Can't access from other devices**
- Ensure devices are on the same network
- Check firewall settings on the host machine
- Verify the server is bound to `0.0.0.0` (default)
- Try accessing via the displayed Network URL

**File upload fails**
- Check available disk space in upload directory
- Verify file doesn't exceed size limit (default: 100MB)
- Ensure upload directory has write permissions

**"File too large" error**
```bash
./shareIO --max-file-size 500  # Increase to 500MB
```

**Server fails to start**
- Check if the port is available
- Verify directory permissions for upload path
- Run `./shareIO --help` to verify command syntax

### Debug Information

The application provides helpful startup information:
- Upload directory location
- Current configuration settings
- Access URLs for local and network connections
- File auto-delete timing

### Network Connectivity

To verify network access:
1. Note the "Network" URL displayed at startup
2. Try accessing from another device's browser
3. Check that devices are on the same network subnet
4. Temporarily disable firewalls for testing

## Use Cases

- **Development Teams**: Quick file sharing during meetings
- **Presentations**: Share slides and documents in conference rooms
- **Home Networks**: Transfer files between family devices
- **Offline Environments**: File sharing without internet access
- **Temporary Collaboration**: Short-term file exchange with auto-cleanup

## License

This project is open source. See LICENSE file for details.

## Contributing

Contributions welcome! Please submit issues and pull requests on GitHub.

---

**Note**: This tool is designed for trusted local networks. Always consider your network security when sharing files.
