package usecase

import (
	"agahi/internal/delivery/dto"
	"agahi/internal/entity/users"
	"agahi/internal/platform/utilities"
)

func Register(r users.Repository, req dto.RegisterUserRequest) error {
	user := users.User{
		Number:   req.Number,
		Email:    req.Email,
		Password: utilities.PasswordHash(req.Password),
	}
	err := r.RegisterUser(user)
	if err != nil {
		return err
	}
	return nil
}
