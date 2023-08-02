package run

import (
	"net/http"

	"github.com/EugeneGpil/router"
	"github.com/r3labs/sse/v2"

	sseRouter "github.com/EugeneGpil/golang_sse/app/router"
)

func Run() {
	server := sse.New()

	server.CreateStream("messages")

	sseRouter.Register(server)

	mux := http.NewServeMux()

	router.DefineRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
