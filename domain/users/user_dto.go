package users

import (
	"strings"

	"github.com/suvamsingh/bookstore_users-api/utils/errors"
)

type (
	//User ...
	User struct {
		ID          int64  `json:"id,omitempty"`
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Email       string `json:"email,omitempty"`
		CreatedDate string `json:"created_date,omitempty"`
	}
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
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
