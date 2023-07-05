package main

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/routes/get_events"
	"github.com/EugeneGpil/golang_sse/app/routes/post_message"
	"github.com/EugeneGpil/golang_sse/app/routes/post_remove_stream"
	"github.com/EugeneGpil/router"

	"github.com/EugeneGpil/golang_sse/app/globals"
)

func main() {
	sever := globals.GetServer()

	sever.CreateStream("messages")

	get_events.Run()
	post_message.Run()
	post_remove_stream.Run()

	mux := http.NewServeMux()

	router.DefineRoutes(mux)
	
	http.ListenAndServe(":8080", mux)
}
