package events

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals"
	"github.com/EugeneGpil/golang_sse/app/utils/defineRoute"
)

func Define() {
	server := globals.GetServer()

	defineRoute.Run(http.MethodGet, "/events/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	})
}
