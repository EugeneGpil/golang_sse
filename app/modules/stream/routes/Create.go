package routes

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/controllers"
	"github.com/EugeneGpil/router"
)

func Create() {
	router.AddRoute(http.MethodPost, "stream", controllers.Create)
}
