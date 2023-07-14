package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	db         = "cinnox"
	collection = "users"
)

type User struct {
	ID       string
	Username string
	Email    string
}

func (m *MongoStore) InsertUser(client *mongo.Client, user User) error {
	// collection := client.Database(db).Collection(collection)
	insertResult, err := m.connPool.Database(db).Collection(collection).InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted user with ID: %s\n", insertResult.InsertedID)
	return nil
}

func (m *MongoStore) FindUser(client *mongo.Client, username string) (User, error) {
	var user User
	// collection := client.Database(db).Collection(collection)
	filter := bson.M{"username": username}
	err := m.connPool.Database(db).Collection(collection).FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m *MongoStore) UpdateUserEmail(client *mongo.Client, username string, newEmail string) error {
	// collection := client.Database(db).Collection(collection)
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"email": newEmail}}
	updateResult, err := m.connPool.Database(db).Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("Updated %d user(s)\n", updateResult.ModifiedCount)
	return nil
}

func (m *MongoStore) DeleteUser(client *mongo.Client, username string) error {
	// collection := client.Database(db).Collection(collection)
	filter := bson.M{"username": username}
	deleteResult, err := m.connPool.Database(db).Collection(collection).DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %d user(s)\n", deleteResult.DeletedCount)
	return nil
}
