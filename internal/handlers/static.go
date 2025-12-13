package handlers

import (
	"net/http"
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
	// Clean the path to resolve any ".." sequences
	cleanPath := filepath.Clean(path)

	// Ensure the path doesn't go outside the static directory
	if !strings.HasPrefix(cleanPath, "static/") && cleanPath != "static" {
		c.String(http.StatusBadRequest, "Invalid path")
		return
	}

	// Additional check: ensure no directory traversal after cleaning
	if strings.Contains(cleanPath, "..") {
		c.String(http.StatusBadRequest, "Invalid path")
		return
	}

	fullPath := filepath.Join(".", path)

	// Handle CSS files - check if CSS exists, otherwise try SCSS
	if strings.HasSuffix(path, ".css") {
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			// Try SCSS file instead
			scssPath := strings.TrimSuffix(fullPath, ".css") + ".scss"
			if _, err := os.Stat(scssPath); err == nil {
				fullPath = scssPath
				c.Header("Content-Type", "text/css")
			}
		}
	}

	// Check if file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		c.String(http.StatusNotFound, "File not found: %s", filepathParam)
		return
	}

	// Serve the file
	c.File(fullPath)
}
