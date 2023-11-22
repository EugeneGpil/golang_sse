package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/validators"
	"github.com/EugeneGpil/golang_sse/app/ship/response"
	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
)

func Create(writer http.ResponseWriter, request *http.Request) {
	stream := validators.ValidateStream(writer, request)
	if (stream == "") {
		return
	}

	sseServer.Get().CreateStream(stream)

	response.WriteSuccess(writer)
}
