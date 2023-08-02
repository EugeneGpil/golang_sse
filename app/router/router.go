package router

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals/sseServer"
	"github.com/EugeneGpil/golang_sse/app/router/routes"
	"github.com/EugeneGpil/router"
)

func RegisterRoutes(mux *http.ServeMux) {
	server := sseServer.Get()

	server.CreateStream("messages")

	routes.GetEvents()
	routes.RemoveStream()
	routes.SendMessage()

	router.DefineRoutes(mux)
}
