package api

import (
	"encoding/json"
	"net/http"

	"otus/socnet/core"
	"time"
)

func logout(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := request.Cookie("session_token")
	if err != nil {
		message := core.ResponseError(err)
		json, _ := json.Marshal(message)
		writeResponse(writer, json)
		return
	}

	sessionToken := cookie.Value
	delete(sessions, sessionToken)
	deleteCookie(writer)

	message := core.ResponseOK()
	json, _ := json.Marshal(message)
	writeResponse(writer, json)
}

func deleteCookie(writer http.ResponseWriter) {
	http.SetCookie(writer, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}
