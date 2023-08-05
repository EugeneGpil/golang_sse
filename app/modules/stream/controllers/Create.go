package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
)

func Create(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	streamName := query.Get("stream")

	sseServer.Get().CreateStream(streamName)
}
