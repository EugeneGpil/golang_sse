package routes

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals/sseServer"
	"github.com/EugeneGpil/router"
)

func GetEvents() {
	router.AddRoute(http.MethodGet, "/events/", func(w http.ResponseWriter, r *http.Request) {
		sseServer.Get().ServeHTTP(w, r)
	})
}
