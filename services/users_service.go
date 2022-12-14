package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/date_utils"
	"bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDbFormat()
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

func UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.Email = user.Email
		current.LastName = user.LastName
	}

	err = current.Update()
	if err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(id int64) *errors.RestErr {
	user := &users.User{Id: id}

	return user.Delete()
}

func Search(status string) ([]users.User, *errors.RestErr) {

	dao := &users.User{}

	usersWithStatus, err := dao.FindByStatus(status)
	if err != nil {
		return nil, err
	}

	return usersWithStatus, nil
}
