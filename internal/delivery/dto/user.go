package dto

type RegisterUserRequest struct {
	Number   string `bson:"number"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type LoginUserRequest struct {
	Number   string `bson:"number"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
