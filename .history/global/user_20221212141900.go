package global

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID primitive.ObjectID `bson:`
	Username string
	Email string
	Password string
}