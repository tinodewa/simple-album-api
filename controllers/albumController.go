package controllers

import (
	"context"
	"hit/album-mongo-api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	// Connect to MongoDB
	client, err := database.GetMongoClient()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	ctx := context.Background()

	//use the global client
	albums, err := database.GetAlbums(ctx, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}
