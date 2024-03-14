package controllers

import (
	"github.com/EugeneGpil/golang_sse/app/modules/stream/validators/StreamValidator"
	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
	"github.com/EugeneGpil/request"
	"github.com/EugeneGpil/response"
	"github.com/r3labs/sse/v2"
)

func Send(response response.Response, request *request.Request) {
	stream := StreamValidator.New(response, request).Run()

	message := request.GetSimpleQueryParams()["message"]

	sseServer.Get().Publish(stream, &sse.Event{
		Data: []byte(message),
	})
}
