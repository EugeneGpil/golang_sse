package router

import (
	"net/http"

	"github.com/EugeneGpil/getFormattedUrls"
)

func DefineRoutes () {
	urlRoutesMap := make(map[string][]Route)

	for _, route := range routes {
		if _, isset := urlRoutesMap[route.url]; !isset {
			urlRoutesMap[route.url] = []Route{}
		}
		urlRoutesMap[route.url] = append(urlRoutesMap[route.url], route)
	}

	for url, routes := range urlRoutesMap {
		url1, url2 := getFormattedUrls.Run(url)

		addHandlers(url1, routes)
		addHandlers(url2, routes)
	}
}

func addHandlers(url string, routes []Route) {
	m.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		for _, route := range routes {
			if request.Method == route.method {
				route.handler(writer, request)

				return
			}
		}

		writer.WriteHeader(http.StatusNotFound)
	})
}
