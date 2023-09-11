package repository

import (
	"agahi/internal/entity/users"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	DB *mongo.Database
}

func (r UserRepo) RegisterUser(u users.User) error {
	coll := r.DB.Collection("users")
	_, err := coll.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	return nil
}
