package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(id int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: id}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}
