package defineRoute

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals"
	"github.com/EugeneGpil/golang_sse/app/utils/getFormattedUrls"
)

func Run(httpMethod string, url string, handler func(http.ResponseWriter, *http.Request)) {
	url1, url2 := getFormattedUrls.Run(url)

	addRoute(httpMethod, url1, handler)
	addRoute(httpMethod, url2, handler)
}

func addRoute(method string, url string, handler func(http.ResponseWriter, *http.Request)) {
	mux := globals.GetMux()

	mux.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != method {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		handler(writer, request)
	})
}
