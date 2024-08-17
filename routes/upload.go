package routes

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kunniii/image_server/utils"
)

// JSON request struct
type ImageUploadRequest struct {
	Base64 string `json:"base64" binding:"required"`
	Ext    string `json:"ext" binding:"required"`
}

func RegisterUploadRoute(r *gin.Engine) {
	r.POST("/upload", utils.ImageUploadValidator, func(c *gin.Context) {

		base64Image := c.GetString("base64")
		ext := c.GetString("extension")

		imageData, err := base64.StdEncoding.DecodeString(base64Image)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid base64 data"})
			return
		}

		imageUUID := uuid.New().String()
		tempPath := filepath.Join("images", imageUUID+"."+ext)

		if err := os.WriteFile(tempPath, imageData, 0644); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image"})
			return
		}

		hash, err := utils.ComputeFileHash(tempPath)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to compute file hash"})
			return
		}

		finalPath := filepath.Join("images", hash+"."+ext)

		if err := os.Rename(tempPath, finalPath); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to rename file"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"url": "/images/" + hash + "." + ext})
	})
}
