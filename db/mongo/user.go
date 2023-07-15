package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID       string
	Username string
	Email    string
}

func (m *MongoStore) InsertUser(user User) error {
	_, err := m.connPool.Database(db).Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoStore) FindUser(username string) (User, error) {
	var user User
	filter := bson.M{"username": username}
	err := m.connPool.Database(db).Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m *MongoStore) UpdateUserEmail(username string, newEmail string) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"email": newEmail}}
	updateResult, err := m.connPool.Database(db).Collection("users").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("Updated %d user(s)\n", updateResult.ModifiedCount)
	return nil
}

func (m *MongoStore) DeleteUser(username string) error {
	filter := bson.M{"username": username}
	deleteResult, err := m.connPool.Database(db).Collection("users").DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %d user(s)\n", deleteResult.DeletedCount)
	return nil
}
