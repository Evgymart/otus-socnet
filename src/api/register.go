package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"otus/socnet/models"
)

func register(writer http.ResponseWriter, request *http.Request) {
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

	user, err := models.BuildUser(data)
	if err != nil {
		log.Fatalln(err)
		return
	}

	user.FirstName += " Anime girl"
	json, _ := json.Marshal(user)
	writeResponse(writer, json)
}
