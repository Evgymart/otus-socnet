package api

import (
	"encoding/json"
	"io"
	"net/http"
	"otus/socnet/core"
)

func createPost(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if checkLogin(request) == "" {
		http.Error(writer, "Not authorized", http.StatusUnauthorized)
		return
	}

	id, err := getUserId(request)
	if err != nil {
		http.Error(writer, "Error during post creation", http.StatusInternalServerError)
		return
	}

	text := getText(request)
	var message core.Message
	err = core.CreatePost(id, text)
	if err != nil {
		message = core.ResponseError(err)
	} else {
		message = core.ResponseOK()
	}

	json, _ := json.Marshal(message)
	writeResponse(writer, json)
}

func getText(request *http.Request) string {
	var createPostData core.CreatePostData
	defer request.Body.Close()
	data, err := io.ReadAll(request.Body)
	if err != nil {
		return ""
	}

	_ = json.Unmarshal(data, &createPostData)
	return createPostData.Text
}
