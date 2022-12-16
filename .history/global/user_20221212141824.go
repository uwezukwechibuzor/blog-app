package global

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID primitive.ObjectID
	Username string
	Email string
	Password type SortBy []Type
	
	func (a SortBy) Len() int           { return len(a) }
	func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
	func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }
}