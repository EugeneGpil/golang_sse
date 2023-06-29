package router

import (
	"net/http"
)

var m *http.ServeMux
var routes []Route

func SetMux(mux *http.ServeMux) {
	m = mux
}
