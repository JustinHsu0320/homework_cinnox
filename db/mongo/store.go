package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db = "cinnox"
)

// MongoStore provides all functions to execute queries
type MongoStore struct {
	connPool *mongo.Client
}

// NewStore creates a new store
func NewStore(pool *mongo.Client) *MongoStore {
	return &MongoStore{
		connPool: pool,
	}
}
