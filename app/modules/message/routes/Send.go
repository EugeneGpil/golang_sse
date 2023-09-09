package routes

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/modules/message/controllers"
	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/router"
)

func Send() {
	router.AddRoute(http.MethodPost, "message", controllers.Send, names.MessageSend)
}
