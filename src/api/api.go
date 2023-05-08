package api

import (
	"net/http"
	"otus/socnet/core"
)

var sessions = map[string]core.Session{}

func InitApi(mux *http.ServeMux) {
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/user/register", register)
	mux.HandleFunc("/user/get", getUser)
	mux.HandleFunc("/user/search", userSearch)
	mux.HandleFunc("/user/swap_names", swapNames)
}

func writeResponse(writer http.ResponseWriter, responseMessage []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseMessage)
}

func checkLogin(request *http.Request) string {
	cookie, err := request.Cookie("session_token")
	if err != nil {
		return ""
	}

	sessionToken := cookie.Value
	userSession, exists := sessions[sessionToken]
	if !exists {
		return ""
	}

	if userSession.IsExpired() {
		delete(sessions, sessionToken)
		return ""
	}

	return userSession.Email
}
