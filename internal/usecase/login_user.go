package usecase

import (
	"agahi/internal/delivery/dto"
	"agahi/internal/entity/users"
)

func Login(r users.Repository, req dto.LoginUserRequest) error {
	user := users.User{
		Email:    req.Email,
		Password: req.Password,
	}
	_, err := r.LoginUser(user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
