package router

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/modules/message"
	"github.com/EugeneGpil/golang_sse/app/modules/stream"
	"github.com/EugeneGpil/router"
)

func Register(mux *http.ServeMux) {
	message.RegisterRoutes()
	stream.RegisterRoutes()

	router.DefineRoutes(mux)
}
