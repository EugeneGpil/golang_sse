package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals/sseServer"
)

func Listen(w http.ResponseWriter, r *http.Request) {
	sseServer.Get().ServeHTTP(w, r)
}
