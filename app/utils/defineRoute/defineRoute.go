package defineRoute

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals"
)

func Run(httpMethod string, url string, handler func(http.ResponseWriter, *http.Request)) {
	mux := globals.GetMux()

	mux.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		if (request.Method != httpMethod) {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		handler(writer, request)
	})
}
