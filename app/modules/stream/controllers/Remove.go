package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/validators/StreamValidator"
	"github.com/EugeneGpil/golang_sse/app/ship/response"
	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
)

func Remove(writer http.ResponseWriter, request *http.Request) {
	stream := StreamValidator.New(writer, request).Run()
	if (stream == "") {
		return
	}

	sseServer.Get().RemoveStream(stream)

	response.WriteSuccess(writer)
}
