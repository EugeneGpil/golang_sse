package get_events

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals"
	"github.com/EugeneGpil/golang_sse/app/utils/defineRoute"
)

func Run() {
	server := globals.GetServer()
	mux := globals.GetMux()

	defineRoute.Run(mux, http.MethodGet, "events/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	})
}
