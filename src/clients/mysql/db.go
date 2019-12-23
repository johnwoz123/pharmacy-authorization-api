package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_DbUsername = "admin"
	mysql_DbPassword = "password"
	mysql_DbHost     = "localhost:3007"
	mysql_Schema     = "auth_db"
)

var (
	Client          *sql.DB
	mysqlDbUsername = os.Getenv(mysql_DbUsername)
	mysqlDbPassword = os.Getenv(mysql_DbPassword)
	mysqlDbHost     = os.Getenv(mysql_DbHost)
	mysqlSchema     = os.Getenv(mysql_Schema)
)

func init() {
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"password",
		"127.0.0.1:3307",
		"auth_db")
	var err error

	Client, err = sql.Open("mysql", datasource)
	if err != nil {
		log.Fatal(err)
	}

	if err = Client.Ping(); err != nil {
		log.Println(err)
	}
	log.Println("connected to database ........ ")

}
