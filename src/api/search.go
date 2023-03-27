package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"otus/socnet/core"
)

func userSearch(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if checkLogin(request) == "" {
		http.Error(writer, "Not authorized", http.StatusUnauthorized)
		return
	}

	name := getName(request)
	if name == "" {
		message := core.ResponseError(errors.New("No name provided"))
		json, _ := json.Marshal(message)
		writeResponse(writer, json)
		return
	}

	users, err := core.SearchUsers(name)
	if err != nil {
		message := core.ResponseError(err)
		json, _ := json.Marshal(message)
		writeResponse(writer, json)
		return
	}

	data, _ := json.Marshal(users)
	message := core.ResponseData(data)
	json, _ := json.Marshal(message)
	writeResponse(writer, json)
}

func getName(request *http.Request) string {
	var searchData core.SearchData
	defer request.Body.Close()
	data, err := io.ReadAll(request.Body)
	if err != nil {
		return ""
	}

	_ = json.Unmarshal(data, &searchData)
	return searchData.Name
}
