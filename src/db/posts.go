package db

func AddPost(db Database, userId int, text string) error {
	insert := "INSERT INTO posts (user_id, text) VALUES (?, ?)"
	statement, err := db.Client.Query(insert, userId, text)
	if err != nil {
		return err
	}

	defer statement.Close()
	return nil
}
