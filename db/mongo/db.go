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

func GetMongoDBClientPool(ctx context.Context, url string) (*mongo.Client, error) {

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(url).
		SetMaxPoolSize(maxPoolSize).
		SetMinPoolSize(minPoolSize)

	// Create the MongoDB client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}
