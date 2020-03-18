package users

import (
	"strings"

	"github.com/suvamsingh/bookstore_users-api/utils/errors"
)

const (
	//StatusActive status of the user is active
	StatusActive = "active"
)

type (
	//User ...
	User struct {
		ID          int64  `json:"id,omitempty"`
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Email       string `json:"email,omitempty"`
		CreatedDate string `json:"created_date,omitempty"`
		Status      string `json:"status,omitempty"`
		Password    string `json:"password,omitempty"` //we do _ because we dont want
		//to make json with password we will see it in our code later
	}
)

type (
	//Users which we will use for marshalling in or public or private request
	Users []User
)

// //Validate ...
// func Validate(user *User) *errors.RestErr {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("Invalid Email Address")
// 	}
// 	return nil
// }

//instead of adding validation as function above we can declare it ads method

//Validate ...
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid Password")
	}
	return nil
}
