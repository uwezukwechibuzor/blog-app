package global

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID primitive.ObjectID `bso`
	Username string
	Email string
	Password string
}