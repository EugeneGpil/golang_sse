package routes

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/controllers"
	"github.com/EugeneGpil/router"
)

func CreateStream() {
	router.AddRoute(http.MethodPost, "/create_stream/", controllers.CreateStream)
}
