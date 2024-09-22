package main

import (
	"context"
	"hit/album-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

// A Album Store Resource allows you to insert album into collection

func main() {
	// Instantiate Gin
	router := gin.Default()

	// Connect to MongoDB
	client, err := controllers.GetMongoClient()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	albums, err := controllers.GetAlbums(context.Background(), client)
	if err != nil {
		panic(err)
	}

	for _, album := range albums {
		println(album.Title)
	}

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
