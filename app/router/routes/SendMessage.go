package routes

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/controllers"
	"github.com/EugeneGpil/router"
)

func SendMessage() {
	router.AddRoute(http.MethodPost, "/message/", controllers.SendMessage)
}
