package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db = "cinnox"
)

// SQLStore provides all functions to execute SQL queries and transactions
type MongoStore struct {
	connPool *mongo.Client
}

// NewStore creates a new store
func NewStore(pool *mongo.Client) *MongoStore {
	return &MongoStore{
		connPool: pool,
	}
}
