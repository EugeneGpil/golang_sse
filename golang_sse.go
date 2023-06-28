package main

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/routes/events"
	"github.com/EugeneGpil/golang_sse/app/routes/remove_stream"
	"github.com/EugeneGpil/golang_sse/app/routes/message"
	"github.com/r3labs/sse/v2"
)

func main() {
	server := sse.New()
	server.CreateStream("messages")
	mux := http.NewServeMux()

	events.Define(server, mux)
	message.Define(server, mux)
	remove_stream.Define(server, mux)

	http.ListenAndServe(":8080", mux)
}
