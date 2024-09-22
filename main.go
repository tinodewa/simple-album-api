package main

import (
	"hit/album-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func Main() {
	main()
}

func main() {
	// Instantiate Gin
	router := gin.Default()

	// Create a route with GET method
	router.GET("/", func(ctx *gin.Context) {
		// Return JSON response
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// Create a route to get all albums
	router.GET("/api/albums", controllers.GetAlbums)

	// Create server with port 3000
	router.Run(":3000")

}
