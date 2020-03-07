package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //we need this as we are using mysql databse
	//but tis all implements database/sql we rae not using  github.com/go-sql-driver/mysql
	//but we need else it will not recognize mysql
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	//Client instance of sql database
	Client   *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database Successfully Configured ")
}
