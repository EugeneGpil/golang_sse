package main

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals"
	"github.com/EugeneGpil/golang_sse/app/routes/events"
	"github.com/EugeneGpil/golang_sse/app/routes/message"
	"github.com/EugeneGpil/golang_sse/app/routes/remove_stream"
)

func main() {
	mux := globals.GetMux()

	sever := globals.GetServer()

	sever.CreateStream("messages")

	events.Define()
	message.Define()
	remove_stream.Define()

	http.ListenAndServe(":8080", mux)
}
