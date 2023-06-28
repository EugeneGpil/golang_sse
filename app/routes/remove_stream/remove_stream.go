package remove_stream

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals"
	"github.com/EugeneGpil/golang_sse/app/utils/defineRoute"
)

func Define() {
	server := globals.GetServer()

	defineRoute.Run(http.MethodPost, "/remove_stream/", func(w http.ResponseWriter, r *http.Request) {
		server.RemoveStream("messages")
	})
}
