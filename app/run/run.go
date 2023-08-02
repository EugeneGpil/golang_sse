package run

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/router"
)

func Run() {
	mux := http.NewServeMux()

	router.RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
