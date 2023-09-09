package tests

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/router"
)

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()

	router.Register(mux)

	return mux
}
