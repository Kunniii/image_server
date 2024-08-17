package utils

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ImageUploadValidator is a middleware that validates the incoming image upload request
func ImageUploadValidator(c *gin.Context) {
    var req struct {
        Base64 string `json:"base_64" binding:"required"`
        Ext    string `json:"ext" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
        c.Abort()
        return
    }

    validExtensions := map[string]bool{
        "jpg":  true,
        "jpeg": true,
        "png":  true,
    }

    if _, ok := validExtensions[strings.ToLower(req.Ext)]; !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file extension"})
        c.Abort()
        return
    }

    if _, err := base64.StdEncoding.DecodeString(req.Base64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid base64 data"})
        c.Abort()
        return
    }

    c.Set("base64", req.Base64)
    c.Set("extension", req.Ext)

    c.Next()
}
