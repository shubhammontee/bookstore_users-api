package users

import (
	"fmt"
	"strings"

	"github.com/suvamsingh/bookstore_users-api/datasources/mysql/users_db"
	"github.com/suvamsingh/bookstore_users-api/logger"
	"github.com/suvamsingh/bookstore_users-api/utils/date_utils"
	"github.com/suvamsingh/bookstore_users-api/utils/errors"
	"github.com/suvamsingh/bookstore_users-api/utils/mysql_utils"
)

const (
	errorNowRows          = "no rows in result set"
	indexUniqueEmail      = "email_UNIQUE"
	queryInsertUser       = "INSERT INTO users(first_name,last_name,email,created_date,status,password) VALUES(?,?,?,?,?,?);"
	queryGetUser          = "SELECT id,first_name,last_name,email,created_date,status FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
	queryDeleteUser       = "DELETE from users WHERE id=?;"
	queryFindUserByStatus = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE status=?;"
)

// var (
// 	userDB = make(map[int64]*User)
// )

//FindByStatus ...
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	//if we dont do close we will live open connection to database
	//and we will run out of connection very very fast
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		//we need to pass yhe pointer to the scan function or else we will be
		//passing the cpy of user and the user variable above wont get populated
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedDate, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		if len(results) == 0 {
			return nil, errors.NewNotFoundError(fmt.Sprintf("no user matching the status %s", status))
		}
		results = append(results, user)
	}

	return results, nil
}

//Delete ...
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

//Update ...
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

//Get ...
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("Database Error")
	}
	defer stmt.Close()

	//QueryRow as we only need one row from database
	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedDate, &user.Status); err != nil {
		if strings.Contains(err.Error(), errorNowRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user with userid %d  :  %s", user.ID, err.Error()))
	}
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
	//user.CreatedDate = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedDate, user.Status, user.Password)

	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists ", user.Email))
		}

		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user : %s", err.Error()))
	}
	userID, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user : %s", err.Error()))
	}
	user.ID = userID

	// if userDB[user.ID] != nil {
	// 	if userDB[user.ID].Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", user.ID))
	// }
	// user.CreatedDate = date_utils.GetNowString()
	// userDB[user.ID] = user
	// return nil

	return nil
}

//now as we can see we are checking nested error by strings.contains
//so for that we can use the a cool interface
//we have MYSqlError which implements this interface
//as mysql error gives an error number we will be using this number
//we can check this by printing the sql erroe in above Save function

//The struct is
// type MYSqlError struct{
// 	Number uint16
// 	Message string
// }

//SaveUsingMysqlErrorNumber ...
func (user *User) SaveUsingMysqlErrorNumber() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.CreatedDate = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedDate)

	if saveErr != nil {

		return mysql_utils.ParseError(saveErr)

		// //now here we are saying if saveErr than attemp to convert it to MySQLError
		// sqlErr, ok := saveErr.(*mysql.MySQLError)
		// if !ok {
		// 	//means this saveErr is not type of MySQlError
		// 	return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user : %s", err.Error()))
		//we will throw internal server erroe as we dont any other way of handling this error
	}

	userID, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user : %s", err.Error()))
	}

	user.ID = userID
	return nil

}

//so in the above method what we are saying is if you dont have MySQLError number
//return internel server error
//and if you have the MySQLError number than return error based on that no

//similarly for Get method

//GetUsingMysqlErrorNumber ...
func (user *User) GetUsingMysqlErrorNumber() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		//using logger
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("Database Error")
		//return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//QueryRow as we only need one row from database
	result := stmt.QueryRow(user.ID)

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedDate); getErr != nil {
		//mysql_utils.ParseError(getErr)
		//using logger
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("Database Error")
	}
	return nil

}

//but in our get the getErr cannnot be converted to MySQLError so here
//we don get any error number returned in getErr we can print and see that
//so we cannot use MySQLError.Number to Handle error
