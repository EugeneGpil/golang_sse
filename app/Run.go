package app

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/router"
)

func Run() {
	mux := http.NewServeMux()

	router.Register(mux)

	http.ListenAndServe(":8080", mux)
}
