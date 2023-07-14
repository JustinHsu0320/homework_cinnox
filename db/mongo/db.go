package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	maxPoolSize = 10
	minPoolSize = 1
)

type Queries struct{
	InsertUser(client *mongo.Client, user User) error
	FindUser(client *mongo.Client, username string) (User, error)
	UpdateUserEmail(client *mongo.Client, username string, newEmail string) error
	DeleteUser(client *mongo.Client, username string) error
}

func GetMongoDBClientPool(url string) (*mongo.Client, error) {
	ctx := context.Background()

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(url).
		SetMaxPoolSize(maxPoolSize).
		SetMinPoolSize(minPoolSize)

	// Create the MongoDB client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	return client, nil
}
