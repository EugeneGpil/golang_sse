package routes

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/controllers"
	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/router"
)

func Listen() {
	router.AddRoute(http.MethodGet, "stream", controllers.Listen, names.StreamListen)
}
