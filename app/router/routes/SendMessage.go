package routes

import (
	"net/http"

	"github.com/EugeneGpil/router"
	"github.com/r3labs/sse/v2"
)

func SendMessage(server *sse.Server) {
	router.AddRoute(http.MethodPost, "/message/", func(w http.ResponseWriter, request *http.Request) {
		query := request.URL.Query()

		streamName := query.Get("stream")
		message := query.Get("message")

		server.Publish(streamName, &sse.Event{
			Data: []byte(message),
		})
	})
}
