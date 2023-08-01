package routes

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals"
	"github.com/EugeneGpil/router"
)

func RemoveStream() {
	server := globals.GetServer()

	router.AddRoute(http.MethodPost, "/remove_stream/", func(w http.ResponseWriter, r *http.Request) {
		server.RemoveStream("messages")
	})
}
