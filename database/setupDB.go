package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect app to mongoDB atlas
func GetMongoClient() (*mongo.Client, error) {
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
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("There was a problem connectiong to your Atlas cluster. Check that the URI includes a valid username and password, and that your IP address has been added to the access list. Error: ")
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client, nil
}
