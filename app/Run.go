package app

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/router"
	"github.com/EugeneGpil/golang_sse/app/ship/translator"
)

func Run() {
	mux := http.NewServeMux()

	router.Register(mux)

	translator.Init()

	http.ListenAndServe(":8080", mux)
}
