package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"

	requestPackage "github.com/EugeneGpil/request"
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

	sseServer.Get().CreateStream(body.Stream)

	writer.WriteHeader(http.StatusOK)
}
