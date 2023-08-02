package router

import (
	"github.com/EugeneGpil/golang_sse/app/router/routes"
	"github.com/r3labs/sse/v2"
)

func Register(server *sse.Server) {
	routes.GetEvents(server)
	routes.RemoveStream(server)
	routes.SendMessage(server)
}
