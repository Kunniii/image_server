package main

import (
	"os"

	"github.com/kunniii/image_server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes
	routes.RegisterPingRoute(r)
	routes.RegisterUploadRoute(r)

	// Serve static files in the "images" directory
	r.Static("/images", "./images")

	// Ensure the images directory exists
	os.MkdirAll("images", os.ModePerm)

	// Start the server on port 8080
	r.Run(":8080")
}
