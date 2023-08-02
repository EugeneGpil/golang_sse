package routes

import (
	"net/http"

	"github.com/EugeneGpil/router"
	"github.com/r3labs/sse/v2"
)

func GetEvents(server *sse.Server) {
	router.AddRoute(http.MethodGet, "/events/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	})
}
