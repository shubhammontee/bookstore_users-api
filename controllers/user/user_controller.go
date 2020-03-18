package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suvamsingh/bookstore_users-api/domain/users"
	"github.com/suvamsingh/bookstore_users-api/services"
	"github.com/suvamsingh/bookstore_users-api/utils/errors"
)

//Get ...
func Get(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, getErr := services.UsersService.GetUser(userID)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

//Create ...
func Create(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	return
	// }

	//we can replace the above commented line with just one line of code
	//which will read the request body as well as unmarshall it

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		//TODO:return bad request to the user
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: Handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//SearchUser ...
// func SearchUser(c *gin.Context) {

// }

//Update ...
func Update(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID

	//this is going to be true in case of patch and false in case of put
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//Delete ...
func Delete(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

//Search ...
func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userID, nil
}
