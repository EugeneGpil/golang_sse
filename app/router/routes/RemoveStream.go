package routes

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/controllers"
	"github.com/EugeneGpil/router"
)

func RemoveStream() {
	router.AddRoute(http.MethodDelete, "/remove_stream/", controllers.RemoveStream)
}
