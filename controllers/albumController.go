package controllers

import (
	"context"
	"fmt"
	"hit/album-mongo-api/models"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StoreAndGetAlbum() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	mongoUri := os.Getenv("MONGO_URI")

	// Connect to your atlas cluster:
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoUri,
	))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("There was a problem connectiong to your Atlas cluster. Check that the URI includes a valid username and password, and that your IP address has been added to the access list. Error: ")
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Setting db and collection
	// Provide the name of the database and collection you want to use.
	// If they don't already exist, the driver and Atlas will create them
	// automatically when you first write data.
	var dbName = os.Getenv("DB_NAME")
	var collectionName = os.Getenv("COLLECTION_NAME")
	collection := client.Database(dbName).Collection(collectionName)

	/*      *** INSERT DOCUMENTS ***
	 *
	 * You can insert individual documents using collection.Insert().
	 * In this example, we're going to create 1 documents and then
	 * insert it in one call with InsertOne().
	 */

	darkAlbum := models.Album{
		Title:  "The Dark Side of the Moon",
		Artist: "Pink Floyd",
		Price:  19.99,
	}

	insertResult, err := collection.InsertOne(ctx, darkAlbum)
	if err != nil {
		fmt.Println("Something went wrong trying to insert the new documents:")
		panic(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	/*
	 * *** FIND DOCUMENTS ***
	 *
	 * Now that we have data in Atlas, we can read it. To retrieve all of
	 * the data in a collection, we create a filter for albums and sort by
	 * name (ascending)
	 */

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Something went wrong trying to find the documents")
		panic(err)
	}

	defer func() {
		cursor.Close(context.Background())
	}()

	// Loop through the returned albums
	for cursor.Next(ctx) {
		album := models.Album{}
		err := cursor.Decode(&album)

		// If there is an error decoding the cursor into an Album
		if err != nil {
			fmt.Println("cursor.Next() error:")
			panic(err)
		} else {
			fmt.Println(album.Title, "created by", album.Artist)
		}
	}
}
