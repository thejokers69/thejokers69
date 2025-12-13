package handlers

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// StaticHandler serves static files from the static directory
func StaticHandler(c *gin.Context) {
	// Get the filepath from the route parameter
	filepathParam := c.Param("filepath")

	// Remove leading slash
	filepathParam = strings.TrimPrefix(filepathParam, "/")

	// Map /assets/* to static/*
	path := filepathParam
	if strings.HasPrefix(path, "assets/") {
		path = strings.Replace(path, "assets/", "static/", 1)
	} else if !strings.HasPrefix(path, "static/") {
		path = "static/" + path
	}

	// Security: prevent directory traversal
	// Unescape the path to handle encoded sequences like %2e%2e
	if decoded, err := url.PathUnescape(filepathParam); err == nil && decoded != "" {
		// prefer decoded path if decoding succeeded
		path = strings.TrimPrefix(decoded, "/")
		// Remap assets -> static again (in case decoded changed it)
		if strings.HasPrefix(path, "assets/") {
			path = strings.Replace(path, "assets/", "static/", 1)
		} else if !strings.HasPrefix(path, "static/") {
			path = "static/" + path
		}
	}

	// Clean the path to resolve any ".." sequences
	cleanPath := filepath.Clean(path)

	// Get the absolute path to the static directory
	staticRootAbs, err := filepath.Abs("static")
	if err != nil {
		c.String(http.StatusInternalServerError, "Server error")
		return
	}

	// Resolve the absolute path to the requested file
	fullPath := filepath.Join(".", cleanPath)
	fullAbsPath, err := filepath.Abs(fullPath)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid path")
		return
	}

	// Resolve symlinks to ensure a symlink doesn't make the path escape
	if resolved, err := filepath.EvalSymlinks(fullAbsPath); err == nil {
		fullAbsPath = resolved
	}

	// Ensure the path is strictly within the static directory
	if !strings.HasPrefix(fullAbsPath, staticRootAbs+string(os.PathSeparator)) && fullAbsPath != staticRootAbs {
		c.String(http.StatusBadRequest, "Invalid path")
		return
	}

	// Handle CSS files - check if CSS exists, otherwise try SCSS
	if strings.HasSuffix(cleanPath, ".css") {
		if _, err := os.Stat(fullAbsPath); os.IsNotExist(err) {
			// Try SCSS file instead
			scssPath := strings.TrimSuffix(fullAbsPath, ".css") + ".scss"
			if _, err := os.Stat(scssPath); err == nil {
				fullAbsPath = scssPath
				c.Header("Content-Type", "text/css")
			}
		}
	}

	// Check if file exists
	if _, err := os.Stat(fullAbsPath); os.IsNotExist(err) {
		c.String(http.StatusNotFound, "File not found: %s", filepathParam)
		return
	}

	// Serve the file
	c.File(fullAbsPath)
}
