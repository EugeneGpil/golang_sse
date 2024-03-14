package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/validators/StreamValidator"
	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
	"github.com/EugeneGpil/request"
	"github.com/EugeneGpil/response"
)

func Listen(response response.Response, request *request.Request) {
	stream := StreamValidator.New(response, request).Run()
	if stream == "" {
		return
	}

	if !sseServer.Get().StreamExists(stream) {
		response.SetStatusCode(http.StatusNotFound)

		return
	}

	sseServer.
		Get().
		ServeHTTP(response.GetWriter(), request.GetRequest())
}
