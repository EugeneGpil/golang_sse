package message

import (
	"net/http"

	"github.com/r3labs/sse/v2"
)

func Define(server *sse.Server, mux *http.ServeMux) {
	mux.HandleFunc("/message", func(w http.ResponseWriter, request *http.Request) {
		if (request.Method != http.MethodPost) {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		query := request.URL.Query()

		streamName := query.Get("stream")
		message := query.Get("message")

		server.Publish(streamName, &sse.Event{
			Data: []byte(message),
		})
	})
}
