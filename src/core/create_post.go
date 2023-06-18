package core

import "otus/socnet/db"

type CreatePostData struct {
	Text string `json:"text"`
}

func CreatePost(userId int, text string) error {
	database := db.GetWriteDb()
	return db.AddPost(database, userId, text)
}
