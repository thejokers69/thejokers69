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

	// Unescape the path to handle encoded sequences like %2e%2e
	unescapedPath, err := url.PathUnescape(filepathParam)
	if err == nil && unescapedPath != "" {
		// Use unescaped when available
		path = strings.TrimPrefix(unescapedPath, "/")
	}

	// Clean the path to resolve any ".." sequences
	cleanPath := filepath.Clean(path)

	// Ensure the cleaned path is under the static/ directory
	if !strings.HasPrefix(cleanPath, "static/") && cleanPath != "static" {
		c.String(http.StatusBadRequest, "Invalid path")
		return
	}

	// Compute absolute paths and ensure the requested file is within the static base directory
	baseDir := filepath.Join(".", "static")
	baseDirAbs, err := filepath.Abs(baseDir)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server error")
		return
	}

	// Join the cleaned path to the workspace and resolve any symlinks
	requested := filepath.Join(".", cleanPath)
	requestedAbs, err := filepath.Abs(requested)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid path")
		return
	}
	// Resolve symlinks to prevent escaping via links
	requestedResolved, err := filepath.EvalSymlinks(requestedAbs)
	if err == nil {
		requestedAbs = requestedResolved
	}

	// Ensure the final resolved path is inside baseDirAbs
	if requestedAbs != baseDirAbs && !strings.HasPrefix(requestedAbs, baseDirAbs+string(filepath.Separator)) {
		c.String(http.StatusBadRequest, "Invalid path")
		return
	}

	// Prevent any lingering '..' sequences (defense in depth)
	if strings.Contains(cleanPath, "..") {
		c.String(http.StatusBadRequest, "Invalid path")
		return
	}

	fullPath := requestedAbs

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
