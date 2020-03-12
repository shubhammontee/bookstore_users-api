package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/suvamsingh/bookstore_users-api/utils/errors"
)

const (
	errorNowRows = "no rows in result set"
)

//ParseError ...
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNowRows) {
			return errors.NewNotFoundError("no record matching the given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	//where 1062 number is thrown by MySQLError if we have duplicate entry
	//for column which is defined as unique is table schema
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
