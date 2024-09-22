package database

import (
	"context"
	"hit/album-mongo-api/models"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertAlbum(ctx context.Context, client *mongo.Client, album models.Album) (*mongo.InsertOneResult, error) {
	// Setting db and collection
	// Provide the name of the database and collection you want to use.
	// If they don't already exist, the driver and Atlas will create them
	// automatically when you first write data.
	var dbName = os.Getenv("DB_NAME")
	var collectionName = os.Getenv("COLLECTION_NAME")
	collection := client.Database(dbName).Collection(collectionName)

	insertResult, err := collection.InsertOne(ctx, album)
	if err != nil {
		return nil, err
	}

	return insertResult, nil

}

func GetAlbums(ctx context.Context, client *mongo.Client) ([]models.Album, error) {
	// Setting db and collection
	// Provide the name of the database and collection you want to use.
	// If they don't already exist, the driver and Atlas will create them
	// automatically when you first write data.
	var dbName = os.Getenv("DB_NAME")
	var collectionName = os.Getenv("COLLECTION_NAME")
	collection := client.Database(dbName).Collection(collectionName)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var albums []models.Album
	if err = cursor.All(ctx, &albums); err != nil {
		return nil, err
	}

	return albums, nil
}
