package api

import "net/http"

func login(writer http.ResponseWriter, request *http.Request) {
	var responseMessage []byte = []byte(`{"mesage": "OK"}`)
	writeResponse(writer, responseMessage)
}
