package routes

import (
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/kunniii/image_server/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterUploadRoute(r *gin.Engine) {
    r.POST("/upload", utils.ImageUploadValidator, func(c *gin.Context) {
        file, _ := c.Get("file")
        ext, _ := c.Get("extension")

        // Generate UUID for the image name
        imageUUID := uuid.New().String()
        imagePath := filepath.Join("images", imageUUID+ext.(string))

        // Save the image to disk
        if err := c.SaveUploadedFile(file.(*multipart.FileHeader), imagePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image"})
            return
        }

        // Return the image URL
        c.JSON(http.StatusOK, gin.H{"url": "/images/" + imageUUID + ext.(string)})
    })
}
