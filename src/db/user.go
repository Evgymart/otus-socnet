package db

import (
	"otus/socnet/models"
	"otus/socnet/views"
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

func GetUsers(db Database, limit int) ([]views.User, error) {
	var users []views.User
	selectQuery := "SELECT first_name, last_name, birthdate, gender, email, biography, city FROM users LIMIT ?"
	response, err := db.Client.Query(selectQuery, limit)
	if err != nil {
		return users, err
	}

	defer response.Close()
	for response.Next() {
		var user views.User
		err := response.Scan(&user.FirstName, &user.LastName, &user.Birthdate, &user.Gender, &user.Email, &user.Biography, &user.City)
		if err != nil {
			return []views.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(db Database, email string) (*views.User, error) {
	var user views.User
	selectQuery := "SELECT first_name, last_name, birthdate, gender, email, biography, city FROM users WHERE email = ? LIMIT 1"
	response, err := db.Client.Query(selectQuery, email)
	if err != nil {
		return nil, err
	}

	defer response.Close()
	if !response.Next() {
		return nil, nil
	}

	err = response.Scan(&user.FirstName, &user.LastName, &user.Birthdate, &user.Gender, &user.Email, &user.Biography, &user.City)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func SearchUsers(db Database, firstName string, lastName string) ([]views.User, error) {
	var users []views.User
	selectQuery := `SELECT first_name, last_name, birthdate, gender, email, biography, city FROM users ` +
		`WHERE(first_name LIKE ? AND last_name LIKE ?) ORDER BY id`

	firstName = "%" + firstName + "%"
	lastName = "%" + lastName + "%"
	response, err := db.Client.Query(selectQuery, firstName, lastName)
	if err != nil {
		return nil, err
	}

	defer response.Close()
	for response.Next() {
		var user views.User
		err := response.Scan(&user.FirstName, &user.LastName, &user.Birthdate, &user.Gender, &user.Email, &user.Biography, &user.City)
		if err != nil {
			return []views.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
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
