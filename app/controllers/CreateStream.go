package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals/sseServer"
)

func CreateStream(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	streamName := query.Get("stream")

	sseServer.Get().CreateStream(streamName)
}
