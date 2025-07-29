# ShareIO - Local Network File Sharing

ShareIO is a lightweight, cross-platform file sharing server designed for quick and easy file transfers between devices on the same local network. Perfect for offline environments where you need to share files between computers, phones, tablets, and other devices without internet connectivity.

## Features

- 🚀 **Zero Configuration** - Just run and share
- 🌐 **Web-based Interface** - Works on any device with a browser
- 📱 **Cross-platform** - Windows, macOS, Linux support
- ⚡ **Automatic Cleanup** - Files auto-delete after 5 minutes
- 🔒 **Local Network Only** - No internet required, keeps files private
- 📋 **QR Code Support** - Easy access from mobile devices
- 🗂️ **Multiple File Management** - Upload, download, delete individual or all files

## Quick Start

### For End Users

1. **Download** the binary for your operating system from the releases page
2. **Run** the application:
   - Windows: Double-click `shareio.exe` or run from command prompt
   - macOS/Linux: Open terminal and run `./shareio`
3. **Access** the web interface at `http://localhost:8000`
4. **Share** the URL with other devices on your network: `http://YOUR_IP:8000`

### Basic Usage

1. Open your web browser and go to the displayed URL
2. Click "Choose File" to select files for upload
3. Other devices can access the same URL to download files
4. Files are automatically deleted after 5 minutes
5. Use the "Delete All" button to clean up immediately

## Command Line Options

```bash
# Basic usage
./shareio

# Custom port
./shareio --port 9000

# Bind to localhost only (more secure)
./shareio --bind 127.0.0.1

# Change auto-delete time
./shareio --delete-after 10m

# Set maximum file size (in MB)
./shareio --max-file-size 500

# Show help
./shareio --help

# Show version
./shareio --version
```

### Environment Variables

You can also configure ShareIO using environment variables:

```bash
# Set port
export SHAREIO_PORT=9000
./shareio

# Set upload directory
export SHAREIO_UPLOAD_DIR=/path/to/uploads
./shareio
```

## Examples

### Home Network File Sharing
```bash
# Start server accessible to all devices on network
./shareio --bind 0.0.0.0 --port 8000
```

### Secure Local Development
```bash
# Localhost only access
./shareio --bind 127.0.0.1 --port 3000
```

### Conference Room Presentation
```bash
# Longer file retention for presentations
./shareio --delete-after 2h --max-file-size 1000
```

## Network Access

When ShareIO starts, it will display:
- **Local URL**: `http://127.0.0.1:8000` (for the same device)
- **Network URL**: `http://192.168.x.x:8000` (for other devices on the network)

Share the network URL with other devices to allow them to access the file sharing interface.

## Security Considerations

ShareIO is designed for **trusted local networks**. Consider these points:

- ✅ **Good for**: Home networks, office LANs, conference rooms, offline environments
- ⚠️ **Be careful**: Public WiFi, untrusted networks
- 🔒 **Security features**: 
  - Files auto-delete after 5 minutes
  - No permanent storage
  - Local network only (no internet exposure)
  - Path traversal protection

## File Management

- **Upload**: Click to select files
- **Download**: Click on any file in the list
- **Delete**: Remove individual files or all files at once
- **Auto-cleanup**: Files automatically delete after 5 minutes

## Supported Platforms

- **Windows** (x64)
- **macOS** (Intel & Apple Silicon)
- **Linux** (x64)

## Building from Source

### Prerequisites
- Go 1.19 or later
- Git

### Clone and Build
```bash
git clone https://github.com/vknow360/shareIO.git
cd shareIO
go mod download
go build -o shareio .
```

### Cross-compilation
```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o shareio-windows.exe .

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o shareio-macos-intel .

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o shareio-macos-arm64 .

# Linux
GOOS=linux GOARCH=amd64 go build -o shareio-linux .
```

## Project Structure

```
shareIO/
├── main.go           # Application entry point
├── routes/
│   └── routes.go     # HTTP route definitions
├── handlers/
│   ├── upload.go     # File upload handling
│   ├── download.go   # File download handling
│   └── files.go      # File management (list, delete)
├── utils/
│   ├── network.go    # Network utility functions
│   └── dir.go        # Directory management
└── static/
    ├── index.html    # Web interface
    ├── 404.html      # Error page
    └── icon.png      # Application icon
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Web interface |
| POST | `/upload` | Upload file |
| GET | `/files` | List all files |
| GET | `/download/:filename` | Download specific file |
| DELETE | `/files/:filename` | Delete specific file |
| DELETE | `/files` | Delete all files |

## Troubleshooting

### Common Issues

**Port already in use**
```bash
./shareio --port 8001
```

**Can't access from other devices**
- Check firewall settings
- Ensure devices are on same network
- Try: `./shareio --bind 0.0.0.0`

**Files not uploading**
- Check available disk space
- Verify file size limits
- Ensure upload directory permissions

**Server won't start**
- Check if port is available
- Verify network interface binding
- Run with `--help` to see options

### Getting Help

1. Run `./shareio --help` for command options
2. Check the console output for error messages
3. Verify network connectivity between devices
4. Check firewall and antivirus settings

## License

This project is open source. See LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## Changelog

### v1.0.0
- Initial release
- Basic file upload/download functionality
- Auto-delete after 5 minutes
- Web-based interface
- Cross-platform support
