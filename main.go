package main

import (
	"os"
	"time"

	"github.com/kunniii/image_server/routes"
	"github.com/kunniii/image_server/utils"

	"github.com/gin-gonic/gin"
)

// startCleanupRoutine periodically cleans up old files
func startCleanupRoutine() {
	ticker := time.NewTicker(24 * time.Hour) // Run cleanup every 24 hours
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := utils.CleanUpOldFiles(); err != nil {
				// Log or handle the error
				// For simplicity, we're just printing it here
				println("Error cleaning up old files:", err.Error())
			}
		}
	}
}

func main() {
	r := gin.Default()

	// Register routes
	routes.RegisterPingRoute(r)
	routes.RegisterUploadRoute(r)

	// Serve static files in the "images" directory
	r.Static("/images", "./images")

	// Ensure the images directory exists
	os.MkdirAll("images", os.ModePerm)

	// Start a cleanup routine to remove old files every 24 hours
	go startCleanupRoutine()

	// Start the server on port 8080
	r.Run(":8080")
}
