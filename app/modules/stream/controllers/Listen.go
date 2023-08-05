package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
)

func Listen(writer http.ResponseWriter, request *http.Request) {
	sseServer.Get().ServeHTTP(writer, request)
}
