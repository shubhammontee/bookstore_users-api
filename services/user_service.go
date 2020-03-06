package services

import (
	"fmt"

	"github.com/suvamsingh/bookstore_users-api/domain/users"
	"github.com/suvamsingh/bookstore_users-api/utils/errors"
)

//CreateUser ...
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	//trim the spaces at the begining and the end of the email
	//and convert it into lower case
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil

}

//GetUser ...
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}
	userErr := user.Get()
	if userErr != nil {
		err := errors.NewNotFoundError(fmt.Sprintf("no such user with id:%d exist", userID))
		return nil, err
	}
	return user, nil
}
