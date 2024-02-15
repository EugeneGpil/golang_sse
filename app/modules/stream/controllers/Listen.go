package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/validators/StreamValidator"
	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
)

func Listen(writer http.ResponseWriter, request *http.Request) {
	stream := StreamValidator.New(writer, request).Run()
	if (stream == "") {
		return
	}

	if (!sseServer.Get().StreamExists(stream)) {
		writer.WriteHeader(http.StatusNotFound)

		return
	}

	sseServer.Get().ServeHTTP(writer, request)
}
