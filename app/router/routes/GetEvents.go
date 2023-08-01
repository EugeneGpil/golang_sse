package routes

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals"
	"github.com/EugeneGpil/router"
)

func GetEvents() {
	server := globals.GetServer()

	router.AddRoute(http.MethodGet, "/events/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	})
}
