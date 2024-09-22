package models

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB Singleton Struct
type mongoClient struct {
	client *mongo.Client
	ctx    context.Context
	once   sync.Once
}

var mongoInstance mongoClient

// Connect app to mongoDB atlas
func GetMongoClient() (*mongo.Client, context.Context, error) {
	mongoInstance.once.Do(func() {
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
		mongoInstance.client = client
		mongoInstance.ctx = ctx
		fmt.Println("Connected to MongoDB!")
	})
	return mongoInstance.client, mongoInstance.ctx, nil
}
