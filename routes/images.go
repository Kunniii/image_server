package routes

import (
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

func RegisterImageRoute(r *gin.Engine) {
	r.GET("/images/:hash", func(c *gin.Context) {
		hash := c.Param("hash")
		path := filepath.Join("images", hash)

		file, err := os.Open(path)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
			return
		}
		defer file.Close()

		img, format, err := image.Decode(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode image"})
			return
		}

		imgWidth := img.Bounds().Dx()
		imgHeight := img.Bounds().Dy()

		reqWidth := c.Query("w")
		reqHeight := c.Query("h")

		var width, height int
		if reqWidth != "" {
			width, err = strconv.Atoi(reqWidth)
			if err != nil || width > imgWidth {
				width = imgWidth
			}
		}
		if reqHeight != "" {
			height, err = strconv.Atoi(reqHeight)
			if err != nil || height > imgHeight {
				height = imgHeight
			}
		}

		if width > 0 || height > 0 {
			img = resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
		}

		if strings.ToLower(format) == "png" {
			c.Writer.Header().Set("Content-Type", "image/png")
			png.Encode(c.Writer, img)
		} else {
			c.Writer.Header().Set("Content-Type", "image/jpeg")
			jpeg.Encode(c.Writer, img, nil)
		}
	})
}
