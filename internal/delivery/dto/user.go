package dto

type RegisterUserRequest struct {
	Number string `bson:"number"`
	Email  string `bson:"email"`
}
