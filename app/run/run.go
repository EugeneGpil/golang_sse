package run

import (
	"net/http"

	sseRouter "github.com/EugeneGpil/golang_sse/app/router"
	"github.com/EugeneGpil/router"

	"github.com/EugeneGpil/golang_sse/app/globals"
)

func Run() {
	sever := globals.GetServer()

	sever.CreateStream("messages")

	sseRouter.Register()

	mux := http.NewServeMux()

	router.DefineRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
