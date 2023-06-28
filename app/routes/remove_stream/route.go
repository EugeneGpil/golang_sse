package remove_stream

import (
	"net/http"

	"github.com/r3labs/sse/v2"
)

func Define(server *sse.Server, mux *http.ServeMux) {
	mux.HandleFunc("/remove_stream", func(w http.ResponseWriter, r *http.Request) {
		server.RemoveStream("messages")
	})
}
