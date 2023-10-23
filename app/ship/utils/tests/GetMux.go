package tests

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/router"
	"github.com/EugeneGpil/golang_sse/app/ship/translator"
)

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()

	router.Register(mux)

	translator.Init()

	return mux
}
