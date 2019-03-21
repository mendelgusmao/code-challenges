package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ObjectID *primitive.ObjectID `bson:"_id" json:"_id"`
	ID       string              `bson:"id" json:"id"`
	Name     string              `bson:"name" json:"name"`
	Email    string              `bson:"email" json:"email"`
	Password string              `bson:"password" json:"password"`
}
