package routes

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kunniii/image_server/utils"
)

func RegisterUploadRoute(r *gin.Engine) {
	r.POST("/upload", utils.ImageUploadValidator, func(c *gin.Context) {
		file, _ := c.Get("file")
		ext, _ := c.Get("extension")

		// Generate a UUID for the image name
		imageUUID := uuid.New().String()
		tempPath := filepath.Join("images", imageUUID+ext.(string))

		// Save the image to a temporary file
		if err := c.SaveUploadedFile(file.(*multipart.FileHeader), tempPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image"})
			return
		}

		// Compute the hash of the image file
		hash, err := utils.ComputeFileHash(tempPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to compute file hash"})
			return
		}

		// Construct the final path using the hash
		finalPath := filepath.Join("images", hash+ext.(string))

		// Rename the file to use the hash as its name
		if err := os.Rename(tempPath, finalPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to rename file"})
			return
		}

		// Return the image URL
		c.JSON(http.StatusOK, gin.H{"url": "/images/" + hash + ext.(string)})
	})
}
