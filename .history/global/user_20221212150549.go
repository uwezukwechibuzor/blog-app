package global

import "go.mongodb.org/mongo-driver/bson/primitive"

//NilUser is the nil value for user
var NilUser User

//User is the default user struct
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

//GetToken returns the user's 
func (u User) GetToken() string {

}