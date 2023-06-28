package defineRoute

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/globals"
)

func Run(httpMethod string, url string, handler func(http.ResponseWriter, *http.Request)) {
	var lastUrlChar byte
	if len(url) == 0 {
		lastUrlChar = '#'
	} else {
		lastUrlChar = url[len(url)-1]
	}

	urlWithoutLastSlash := url
	if lastUrlChar == '/' {
		urlWithoutLastSlash = urlWithoutLastSlash[:len(urlWithoutLastSlash)-1]
	}

	urlWithSlash := urlWithoutLastSlash + "/"

	addRoute(httpMethod, urlWithSlash, handler)
	addRoute(httpMethod, urlWithoutLastSlash, handler)
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
