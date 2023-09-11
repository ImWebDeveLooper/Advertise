package users

type Repository interface {
	RegisterUser(User) error
}
