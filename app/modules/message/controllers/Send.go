package controllers

import (
	"net/http"

	// "github.com/EugeneGpil/golang_sse/app/modules/stream/validators/StreamValidator"
	// "github.com/EugeneGpil/golang_sse/app/ship/sseServer"
	// "github.com/r3labs/sse/v2"
)

func Send(writer http.ResponseWriter, request *http.Request) {
	// stream := StreamValidator.New(writer, request).Run()
	// message := query.Get("message")

	// sseServer.Get().Publish(stream, &sse.Event{
	// 	Data: []byte(message),
	// })
}
