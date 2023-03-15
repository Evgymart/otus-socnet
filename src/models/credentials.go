package models

import "encoding/json"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func BuildCreds(credData []byte) (*Credentials, error) {
	var creds Credentials
	err := json.Unmarshal(credData, &creds)
	if err != nil {
		return nil, err
	}

	return &creds, nil
}
