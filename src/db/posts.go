package db

import "otus/socnet/models"

func AddPost(db Database, user *models.User) error {
	insert := "INSERT INTO users (first_name, last_name, birthdate, gender, email, password, biography, city) VALUES (?,?,?,?,?,?,?,?)"
	statement, err := db.Client.Query(insert, user.FirstName, user.LastName, user.Birthdate.ToString(), user.Gender, user.Email, user.Password, user.Biography, user.City)
	if err != nil {
		return err
	}

	defer statement.Close()
	return nil
}
