package users

type Repository interface {
	RegisterUser(User) error
	LoginUser(string, string) (*User, error)
}
