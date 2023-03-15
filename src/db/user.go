package db

import (
	"otus/socnet/models"
)

func AddUser(db Database, user *models.User) error {
	insert := "INSERT INTO users (first_name, last_name, birthdate, gender, email, password, biography, city) VALUES (?,?,?,?,?,?,?,?)"
	statement, err := db.Client.Query(insert, user.FirstName, user.LastName, user.Birthdate.ToString(), user.Gender, user.Email, user.Password, user.Biography, user.City)
	if err != nil {
		return err
	}
	defer statement.Close()
	return nil
}

func Login(db Database, creds *models.Credentials) (string, error) {
	var password string
	selectQuery := "SELECT password FROM users WHERE email = ?"
	response, err := db.Client.Query(selectQuery, creds.Email)
	if err != nil {
		return "", err
	}

	defer response.Close()
	if !response.Next() {
		return "", nil
	}

	err = response.Scan(&password)
	if err != nil {
		panic(err)
	}

	return password, nil
}
