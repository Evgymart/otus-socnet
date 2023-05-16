package core

import "time"

type Session struct {
	Id     int
	Email  string
	Expiry time.Time
}

func (session Session) IsExpired() bool {
	return session.Expiry.Before(time.Now())
}

func MakeSession(id int, email string) Session {
	expiresAt := time.Now().Add(2 * time.Hour)
	return Session{
		Id:     id,
		Email:  email,
		Expiry: expiresAt,
	}
}
