package users

import (
	"fmt"

	"github.com/suvamsingh/bookstore_users-api/datasources/mysql/users_db"
	"github.com/suvamsingh/bookstore_users-api/utils/date_utils"
	"github.com/suvamsingh/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name,last_name,email,created_date) VALUES(?,?,?,?);"
)

var (
	userDB = make(map[int64]*User)
)

//Get ...
func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	//we need to close it because we will have a connection for creating the statement with us
	//golang will have to do for us
	defer stmt.Close()
	user.CreatedDate = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedDate)

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user : %s", err.Error()))
	}
	userID, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user : %s", err.Error()))
	}
	user.ID = userID

	if userDB[user.ID] != nil {
		if userDB[user.ID].Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.ID))
	}
	user.CreatedDate = date_utils.GetNowString()
	userDB[user.ID] = user
	// return nil

	return nil
}
