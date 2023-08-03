package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals/sseServer"
)

func RemoveStream(w http.ResponseWriter, r *http.Request) {
	sseServer.Get().RemoveStream("messages")
}
