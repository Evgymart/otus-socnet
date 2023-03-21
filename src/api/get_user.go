package api

import (
	"encoding/json"
	"io"
	"net/http"
	"otus/socnet/core"
)

func getUser(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if checkLogin(request) == "" {
		http.Error(writer, "Not authorized", http.StatusUnauthorized)
		return
	}

	email := getEmail(request)
	var data []byte
	if email == "" {
		users, err := core.GetUsers(50)
		if err != nil {
			message := core.ResponseError(err)
			json, _ := json.Marshal(message)
			writeResponse(writer, json)
			return
		}
		data, _ = json.Marshal(users)
	} else {
		user, err := core.GetUser(email)
		if err != nil {
			message := core.ResponseError(err)
			json, _ := json.Marshal(message)
			writeResponse(writer, json)
			return
		}
		data, _ = json.Marshal(user)
	}

	message := core.ResponseData(data)
	json, _ := json.Marshal(message)
	writeResponse(writer, json)
}

func getEmail(request *http.Request) string {
	var getUserData core.GetUserData
	defer request.Body.Close()
	data, err := io.ReadAll(request.Body)
	if err != nil {
		return ""
	}

	_ = json.Unmarshal(data, &getUserData)
	return getUserData.Email
}
