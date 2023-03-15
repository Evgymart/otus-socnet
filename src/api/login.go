package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"otus/socnet/core"
	"otus/socnet/models"
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

	message := core.Login(creds)
	json, _ := json.Marshal(message)
	writeResponse(writer, json)
}
