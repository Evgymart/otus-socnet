package api

import (
	"io"
	"log"
	"net/http"
)

func register(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer request.Body.Close()
	b, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	writeResponse(writer, b)
}
