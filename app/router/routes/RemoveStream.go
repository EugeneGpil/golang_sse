package routes

import (
	"net/http"

	"github.com/EugeneGpil/router"
	"github.com/r3labs/sse/v2"
)

func RemoveStream(server *sse.Server) {
	router.AddRoute(http.MethodPost, "/remove_stream/", func(w http.ResponseWriter, r *http.Request) {
		server.RemoveStream("messages")
	})
}
