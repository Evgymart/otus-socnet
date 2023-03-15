package core

import "golang.org/x/crypto/bcrypt"

func hashPassword(pass string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		panic("Could not hash a password")
	}

	return string(hashed)
}

func compareHash(recieved string, stored string) bool {
	return bcrypt.CompareHashAndPassword([]byte(stored), []byte(recieved)) == nil
}
