package should_create_stream

import (
	// "io"
	"net/http"
	"net/url"
	"testing"

	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests"
	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/tester"
)

var mux = tests.GetMux()

func Test_should_create_stream(t *testing.T) {
	tester.SetTester(t)

	route := router.ByName(names.StreamCreate)

	urlObj, err := url.Parse(route.Url)
	tester.AssertNil(err)

	request := http.Request{
		Method: route.Method,
		URL:    urlObj,
		// Body: io.ReadCloser{
		// 	io.Reader,
		// 	io.Closer,
		// },
	}

	handler, _ := mux.Handler(&request)

	writer := tester.GetTestResponseWriter()

	handler.ServeHTTP(writer, &request)

	status := writer.GetStatus()

	tester.AssertSame(http.StatusOK, status)
}
