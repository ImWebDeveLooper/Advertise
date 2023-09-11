package usecase

import (
	"agahi/internal/delivery/dto"
	"agahi/internal/entity/users"
)

func Register(r users.Repository, req dto.RegisterUserRequest) error {
	user := users.User{
		Number: req.Number,
		Email:  req.Email,
	}
	err := r.RegisterUser(user)
	if err != nil {
		return err
	}
	return nil
}
