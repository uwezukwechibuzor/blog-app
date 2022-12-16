package global

import "go.mongodb.org/mongo-driver/bson/primitive"

//NilUser is the 
var NilUser User

//User is the default user struct
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}
