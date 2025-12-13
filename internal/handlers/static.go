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
