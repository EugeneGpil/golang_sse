package should_remove_stream

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
	"github.com/EugeneGpil/golang_sse/app/ship/translator"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/AssertDoesStreamExists"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/AssertResponseMessage"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/GetMux"
	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

var mux = GetMux.GetMux()
var streamName = "the_test_stream"
var route = router.ByName(names.StreamRemove)
var expectedMessage = translator.Translate("success")

func Test_should_remove_stream(t *testing.T) {
	tester.SetTester(t)

	sseServer.Get().CreateStream(streamName)

	isStreamExists := sseServer.Get().StreamExists(streamName)

	tester.AssertSame(isStreamExists, true)

	response := httpTester.Request(httpTester.GetRequestDto{
		Method: route.Method,
		Url:    route.Url,
		Query: map[string]string{
			"stream": streamName,
		},
	}, mux)

	tester.AssertSame(http.StatusOK, response.GetStatus())

	AssertResponseMessage.AssertResponseMessage(response, expectedMessage)

	AssertDoesStreamExists.AssertDoesStreamExists(t, streamName, false)
}
