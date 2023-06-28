package events

import (
	"net/http"

	"github.com/r3labs/sse/v2"
)

func Define(server *sse.Server, mux *http.ServeMux) {
	mux.HandleFunc("/events/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	})
}
