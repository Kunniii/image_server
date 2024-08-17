package utils

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// Allowed file extensions and MIME types
var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}
var allowedMimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/gif":  true,
}

// Validate image content
func validateImageContent(file multipart.File) error {
	_, _, err := image.Decode(file)
	return err
}

// Image upload validation handler
func ImageUploadValidator(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image file found"})
		c.Abort()
		return
	}

	// Validate Content-Type header
	contentType := c.Request.Header.Get("Content-Type")
	if !allowedMimeTypes[contentType] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
		c.Abort()
		return
	}

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file extension"})
		c.Abort()
		return
	}

	// Validate file content
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open file"})
		c.Abort()
		return
	}
	defer openedFile.Close()

	if err := validateImageContent(openedFile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image content"})
		c.Abort()
		return
	}

	c.Set("file", file)
	c.Set("extension", ext)
	c.Next()
}
