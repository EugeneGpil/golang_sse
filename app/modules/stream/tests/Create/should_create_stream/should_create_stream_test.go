package should_create_stream

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests"
	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

var mux = tests.GetMux()
var route = router.ByName(names.StreamCreate)
var streamName = "messages"

func Test_should_create_stream(t *testing.T) {
	tester.SetTester(t)

	urlObj, err := url.Parse(route.Url)
	tester.AssertNil(err)

	requestBodyBytes, err := json.Marshal(struct {
		Stream string
	} {
		Stream: streamName,
	})

	request := http.Request{
		Method: route.Method,
		URL:    urlObj,
		Body: io.NopCloser(bytes.NewReader(requestBodyBytes)),
	}

	handler, _ := mux.Handler(&request)

	writer := httpTester.GetTestResponseWriter()

	handler.ServeHTTP(writer, &request)

	status := writer.GetStatus()

	tester.AssertSame(http.StatusOK, status)

	isStreamExists := sseServer.Get().StreamExists(streamName)

	tester.AssertSame(isStreamExists, true)
}
