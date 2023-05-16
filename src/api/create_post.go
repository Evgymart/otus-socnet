package api

import (
	"encoding/json"
	"fmt"
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
	json := []byte(fmt.Sprintf(`{"id": %d, "Text": "%s"}`, id, text))
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
