package routes

import (
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

func RegisterImageRoute(r *gin.Engine) {
	r.GET("/images/:hash", func(c *gin.Context) {
		hash := c.Param("hash")
		ext := filepath.Ext(hash)
		path := filepath.Join("images", hash)

		file, err := os.Open(path)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
			return
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode image"})
			return
		}

		width, _ := strconv.Atoi(c.Query("w"))
		height, _ := strconv.Atoi(c.Query("h"))

		if width > 0 || height > 0 {
			img = resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
		}

		c.Header("Content-Type", "image/"+ext[1:])

		if ext == ".png" {
			png.Encode(c.Writer, img)
		} else if ext == ".jpg" || ext == ".jpeg" {
			jpeg.Encode(c.Writer, img, nil)
		}
	})
}
