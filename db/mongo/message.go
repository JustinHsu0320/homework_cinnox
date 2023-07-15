package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Message struct {
	UserID  string
	Message string
}

func (m *MongoStore) InsertMessage(message Message) error {
	_, err := m.connPool.Database(db).Collection("messages").InsertOne(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoStore) GetUsers() ([]interface{}, error) {
	ctx := context.Background()
	// Define a filter for query (optional)
	filter := bson.M{}

	// Execute the distinct query to fetch unique emails
	users, err := m.connPool.Database(db).Collection("messages").Distinct(ctx, "userid", filter)
	if err != nil {
		return nil, err
	}

	return users, nil
}
