package models

type Post struct {
	Text   string
	UserId int
}

func CreatePost(userId int, text string) *Post {
	var post Post
	post.UserId = userId
	post.Text = text
	return &post
}
