package api

import (
	"net/http"
)

func InitApi(mux *http.ServeMux) {
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/user/register", register)
}

func writeResponse(writer http.ResponseWriter, responseMessage []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseMessage)
}
