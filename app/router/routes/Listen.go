package routes

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/controllers"
	"github.com/EugeneGpil/router"
)

func Listen() {
	router.AddRoute(http.MethodGet, "/listen/", controllers.Listen)
}
