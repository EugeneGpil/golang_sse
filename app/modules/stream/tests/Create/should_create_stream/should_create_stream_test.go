package should_create_stream

import (
	"net/http"
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

	body := struct {
		Stream string
	}{
		Stream: streamName,
	}

	response := httpTester.Request(httpTester.GetRequestDto{
		Method: route.Method,
		Url:    route.Url,
		Body:   body,
	}, mux)

	tester.AssertSame(http.StatusOK, response.ResponseWriter.GetStatus())

	isStreamExists := sseServer.Get().StreamExists(streamName)

	tester.AssertSame(isStreamExists, true)
}
