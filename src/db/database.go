package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseParams struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

type Database struct {
	Client *sql.DB
}

var (
	database Database
	params   *DatabaseParams
)

func GetDatabase() Database {
	if database.Client == nil {
		panic("Database is not initiated")
	}

	return database
}

func getParams() DatabaseParams {
	return DatabaseParams{
		User:     "root",
		Password: "root",
		Database: "database",
		Host:     "mysql",
		Port:     "3306",
	}
}

func InitDatabase() {
	params := getParams()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", params.User, params.Password, params.Host, params.Port, params.Database))
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	database.Client = db
}
