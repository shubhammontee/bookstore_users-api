package users

import (
	"fmt"

	"github.com/suvamsingh/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

//Get ...
func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedDate = result.CreatedDate
	return nil

}

//Save ...
func (user *User) Save() *errors.RestErr {
	if userDB[user.ID] != nil {
		if userDB[user.ID].Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.ID))
	}
	userDB[user.ID] = user
	return nil
}
