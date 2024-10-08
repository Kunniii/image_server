package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/kunniii/image_server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	// Register routes
	routes.RegisterPingRoute(r)
	routes.RegisterUploadRoute(r)
	routes.RegisterImageRoute(r)

	// Ensure the images directory exists
	os.MkdirAll("images", os.ModePerm)

	// Start the server on port 8080
	r.Run("0.0.0.0:8080")
}
