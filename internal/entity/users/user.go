package users

type User struct {
	Number string `bson:"number"`
	Email  string `bson:"email"`
}
