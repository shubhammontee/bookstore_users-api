package app

import (
	"github.com/suvamsingh/bookstore_users-api/controllers/ping"
	"github.com/suvamsingh/bookstore_users-api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", user.CreateUser)
	router.GET("/users/:user_id", user.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
	router.PUT("/users/:user_id", user.UpdateUser)
	router.PATCH("/users/:user_id", user.UpdateUser)
}
