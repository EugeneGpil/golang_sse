package router

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/router/routes"
	"github.com/EugeneGpil/router"
)

func RegisterRoutes(mux *http.ServeMux) {
	routes.CreateStream()
	routes.Listen()
	routes.RemoveStream()
	routes.SendMessage()

	router.DefineRoutes(mux)
}
