package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
	"github.com/r3labs/sse/v2"
)

func Send(w http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	streamName := query.Get("stream")
	message := query.Get("message")

	sseServer.Get().Publish(streamName, &sse.Event{
		Data: []byte(message),
	})
}
