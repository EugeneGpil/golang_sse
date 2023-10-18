package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"

	requestPackage "github.com/EugeneGpil/request"
	responsePackage "github.com/EugeneGpil/response"
)

type requestBody struct {
	Stream string
}

func Create(writer http.ResponseWriter, request *http.Request) {
	var body requestBody

	err := requestPackage.New(request).DecodeBody(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response := responsePackage.New(writer)

	if body.Stream == "" {
		response.WriteValidationErrors(map[string]string{
			"stream": "Empty string is not allowed",
		})

		return
	}

	sseServer.Get().CreateStream(body.Stream)

	writer.WriteHeader(http.StatusOK)
}
