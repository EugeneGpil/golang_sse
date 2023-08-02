package router

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/router/routes"
	"github.com/EugeneGpil/router"
	"github.com/r3labs/sse/v2"
)

func RegisterRoutes(mux *http.ServeMux) {
	server := sse.New()

	server.CreateStream("messages")

	routes.GetEvents(server)
	routes.RemoveStream(server)
	routes.SendMessage(server)

	router.DefineRoutes(mux)
}
