package globals

import (
	"net/http"

	"github.com/r3labs/sse/v2"
)

var server = sse.New()
var mux = http.NewServeMux()

func GetServer() *sse.Server {
	return server
}

func GetMux() *http.ServeMux {
	return mux
}