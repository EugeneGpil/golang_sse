package controllers

import (
	"encoding/json"
	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
	"net/http"
)

type requestBody struct {
	Stream string
}

func Create(writer http.ResponseWriter, request *http.Request) {
	var body requestBody

	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	sseServer.Get().CreateStream(body.Stream)

	writer.WriteHeader(http.StatusOK)
}
