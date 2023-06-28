package getFormattedUrls

// Format url for "net/http.ServeMux.HandleFunc"
// Both returned urls should be provided to HandleFunc function

// You should use this package to prevent "forget trailing slash" error
// when you trying to access your app route
// For example "http://test.local/some/route" instead of "http://test.local/some/route/"
// and "http://test.local/some/route/" instead of "http://test.local/some/route"

// And use it to prevent "forget first slash" when providing url for HandlerFunc function
// For example "some/route/" instead of "/some/route/"

// Input url may be with or without trailing slash "/"
// Input url may be with or without first slash "/"


// Returns 2 urls

// Both with first slash "/"

// First without trailing slash "/"
// And second with trailing slash "/"

// Examples
// Input "some/route".		Output "/some/route", "/some/route/"
// Input "/some/route".		Output "/some/route", "/some/route/"
// Input "some/route/".		Output "/some/route", "/some/route/"
// Input "/some/route/".	Output "/some/route", "/some/route/"

func Run(url string) (string, string) {
	if len(url) == 0 {
		return "", "/"
	}

	urlWithFirsSlash := addFirstSlash(url)

	urlWithoutLastSlash := removeLastSlash(urlWithFirsSlash)

	urlWithSlash := urlWithoutLastSlash + "/"

	return urlWithoutLastSlash, urlWithSlash
}

func addFirstSlash(url string) string {
	firstChar := url[0]

	if firstChar == '/' {
		return url
	}

	return url + "/"
}

func removeLastSlash(url string) string {
	lastChar := url[len(url)-1]

	if lastChar == '/' {
		return url[:len(url)-1]
	}

	return url
}
