package api

import (
	"encoding/json"
	"io"
	"net/http"
	"otus/socnet/core"
)

func swapNames(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if checkLogin(request) == "" {
		http.Error(writer, "Not authorized", http.StatusUnauthorized)
		return
	}

	swapData := getIds(request)
	err := core.SwapUserNames(swapData)
	var message core.Message
	if err == nil {
		message = core.ResponseOK()
	} else {
		message = core.ResponseError(err)
	}

	json, _ := json.Marshal(message)
	writeResponse(writer, json)
}

func getIds(request *http.Request) *core.SwapUserNamesData {
	var swapData core.SwapUserNamesData
	defer request.Body.Close()
	data, err := io.ReadAll(request.Body)
	if err != nil {
		return nil
	}

	_ = json.Unmarshal(data, &swapData)
	return &swapData
}
