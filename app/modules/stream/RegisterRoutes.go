package stream

import "github.com/EugeneGpil/golang_sse/app/modules/stream/routes"

func RegisterRoutes() {
	routes.Create()
	routes.Listen()
	routes.Remove()
}
