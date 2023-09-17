package repository

import (
	"agahi/internal/entity/users"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

func (r UserRepo) LoginUser(email, password string) (*users.User, error) {
	coll := r.DB.Collection("users")
	filter := bson.M{"email": email}
	var user users.User
	err := coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not Found")
		}
		return nil, err
	}
	pErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if pErr != nil {
		return nil, errors.New("incorrect Password")
	}
	return &user, nil
}
