package users

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type dao struct {
	db *mongo.Database
}

func newDAO(db *mongo.Database) *dao {
	return &dao{db: db}
}

func (d *dao) findByEmailAndPassword(email, password string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"email":    email,
		"password": password,
	}
	user := User{}

	if err := d.db.Collection("users").FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
