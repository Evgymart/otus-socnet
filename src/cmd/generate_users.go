package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type DatabaseParams struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

type User struct {
	FirstName string `fake:"{firstname}"`
	LastName  string `fake:"{lastname}"`
	Birthdate string `fake:"{year}-{month}-{day}" format:"2006-01-02"`
	Gender    string `fake:"{randomstring:[male,female]}"`
	Email     string `fake:"skip"`
	Password  string `fake:"skip"`
	Biography string `fake:"{sentence:3}"`
	City      string `fake:"{city}"`
}

const USER_COUNT int = 1000000

func main() {
	db := initDatabase()
	defer db.Close()
	for counter := 1; counter <= USER_COUNT; counter++ {
		user := buildUser(counter)
		insertUser(user, db)
	}
	fmt.Println("OK")
}

func initDatabase() *sql.DB {
	params := DatabaseParams{
		User:     "root",
		Password: "root",
		Database: "database",
		Host:     "127.0.0.1",
		Port:     "3306",
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", params.User, params.Password, params.Host, params.Port, params.Database))
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func buildUser(counter int) User {
	var user User
	gofakeit.Struct(&user)
	user.Email = strings.ToLower(user.FirstName) + "_" + strings.ToLower(user.LastName) + strconv.Itoa(counter) + "@gmail.com"
	pass, _ := bcrypt.GenerateFromPassword([]byte("qwerty"), 8)
	user.Password = string(pass)
	return user
}

func insertUser(user User, db *sql.DB) error {
	insert := "INSERT INTO users (first_name, last_name, birthdate, gender, email, password, biography, city) VALUES (?,?,?,?,?,?,?,?)"
	statement, err := db.Query(insert, user.FirstName, user.LastName, user.Birthdate, user.Gender, user.Email, user.Password, user.Biography, user.City)
	if err != nil {
		return err
	}
	defer statement.Close()
	return nil
}
