package defineRoute

import (
	"net/http"

	"github.com/EugeneGpil/getFormattedUrls"
)

func Run(mux *http.ServeMux, httpMethod string, url string, handler func(http.ResponseWriter, *http.Request)) {
	url1, url2 := getFormattedUrls.Run(url)

	addRoute(mux, httpMethod, url1, handler)
	addRoute(mux, httpMethod, url2, handler)
}

func addRoute(mux *http.ServeMux, method string, url string, handler func(http.ResponseWriter, *http.Request)) {
	mux.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != method {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		handler(writer, request)
	})
}
