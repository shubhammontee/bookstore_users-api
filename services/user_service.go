package services

import (
	"fmt"

	"github.com/suvamsingh/bookstore_users-api/domain/users"
	"github.com/suvamsingh/bookstore_users-api/utils/date_utils"
	"github.com/suvamsingh/bookstore_users-api/utils/errors"
)

//CreateUser ...
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	//trim the spaces at the begining and the end of the email
	//and convert it into lower case
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.CreatedDate = date_utils.GetNowDBFormat()
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

//UpdateUser ...
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
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
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

//DeleteUser ...
func DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

//Search ...
func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
