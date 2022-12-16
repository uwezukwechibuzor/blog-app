package global

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID primitive.ObjectID `bson:"_id"`
	Username string`bson:"username"`
	Email string `bson:"email"`
	Password string `bson:"password"`
}