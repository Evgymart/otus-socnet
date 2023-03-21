package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"otus/socnet/core"
	"otus/socnet/models"

	"github.com/google/uuid"
)

func login(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer request.Body.Close()
	data, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	creds, err := models.BuildCreds(data)
	if err != nil {
		log.Fatalln(err)
		return
	}

	success := core.Login(creds)
	var message core.Message
	if success {
		makeSession(creds, writer, request)
		message = core.ResponseOK()
	} else {
		message = core.ResponseError(errors.New("Could not log in"))
	}

	json, _ := json.Marshal(message)
	writeResponse(writer, json)
}

func makeSession(creds *models.Credentials, writer http.ResponseWriter, request *http.Request) {
	sessionToken := uuid.NewString()
	sessionData := core.MakeSession(creds.Email)

	sessions[sessionToken] = sessionData
	http.SetCookie(writer, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: sessionData.Expiry,
	})
}
