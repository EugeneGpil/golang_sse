package router

import (
	"github.com/EugeneGpil/golang_sse/app/router/routes"
)

func Register() {
	routes.GetEvents()
	routes.RemoveStream()
	routes.SendMessage()
}
