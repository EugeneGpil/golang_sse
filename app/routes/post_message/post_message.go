package post_message

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals"
	"github.com/EugeneGpil/golang_sse/app/utils/defineRoute"
	"github.com/r3labs/sse/v2"
)

func Run() {
	server := globals.GetServer()

	defineRoute.Run(http.MethodPost, "/message/", func(w http.ResponseWriter, request *http.Request) {
		query := request.URL.Query()

		streamName := query.Get("stream")
		message := query.Get("message")

		server.Publish(streamName, &sse.Event{
			Data: []byte(message),
		})
	})
}
