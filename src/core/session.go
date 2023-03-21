package core

import "time"

type Session struct {
	Email  string
	Expiry time.Time
}

func (session Session) IsExpired() bool {
	return session.Expiry.Before(time.Now())
}

func MakeSession(email string) Session {
	expiresAt := time.Now().Add(2 * time.Hour)
	return Session{
		Email:  email,
		Expiry: expiresAt,
	}
}
