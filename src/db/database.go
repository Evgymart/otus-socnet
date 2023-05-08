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
	readDbs       [2]Database
	readDbCounter int
	writeDb       Database
)

func GetWriteDb() Database {
	if writeDb.Client == nil {
		panic("Database is not initiated")
	}

	return writeDb
}

func GetReadDb() Database {
	if readDbCounter >= len(readDbs) {
		readDbCounter = 0
	}

	db := readDbs[readDbCounter]
	readDbCounter++
	return db
}

func initDatabase(params *DatabaseParams) Database {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", params.User, params.Password, params.Host, params.Port, params.Database))
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return Database{
		Client: db,
	}
}

func Init() {
	master := DatabaseParams{
		User:     "root",
		Password: "root",
		Database: "database",
		Host:     "mysql_master",
		Port:     "3306",
	}

	slave_one := DatabaseParams{
		User:     "root",
		Password: "root",
		Database: "database",
		Host:     "mysql_slave_first",
		Port:     "3306",
	}

	slave_two := DatabaseParams{
		User:     "root",
		Password: "root",
		Database: "database",
		Host:     "mysql_slave_second",
		Port:     "3306",
	}

	writeDb = initDatabase(&master)
	readDbs[0] = initDatabase(&slave_one)
	readDbs[1] = initDatabase(&slave_two)
	readDbCounter = 0
}
