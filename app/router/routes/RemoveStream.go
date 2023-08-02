package routes

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals/sseServer"
	"github.com/EugeneGpil/router"
)

func RemoveStream() {
	router.AddRoute(http.MethodPost, "/remove_stream/", func(w http.ResponseWriter, r *http.Request) {
		sseServer.Get().RemoveStream("messages")
	})
}
