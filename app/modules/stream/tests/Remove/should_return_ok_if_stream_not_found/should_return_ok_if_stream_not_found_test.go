package should_return_ok_if_stream_not_found

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
	"github.com/EugeneGpil/golang_sse/app/ship/translator"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/AssertDoesStreamExists"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/GetMux"
	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

var mux = GetMux.GetMux()
var streamName = "the_test_stream"
var route = router.ByName(names.StreamRemove)
var expectedMessage = translator.Translate("success")

func Test_should_return_ok_if_stream_not_found(t *testing.T) {
	tester.SetTester(t)

	isStreamExists := sseServer.Get().StreamExists(streamName)

	tester.AssertSame(isStreamExists, false)

	response := httpTester.Request(httpTester.GetRequestDto{
		Method: route.Method,
		Url:    route.Url,
		Query: map[string]string{
			"stream": streamName,
		},
	}, mux)

	tester.AssertSame(http.StatusOK, response.GetStatus())

	//TODO tester.assertMessage(response, expectedMessage)
	responseBody := struct {
		Message string
	} {}

	err := response.DecodeBody(&responseBody)

	tester.AssertNil(err)

	tester.AssertSame(expectedMessage, responseBody.Message)

	AssertDoesStreamExists.AssertDoesStreamExists(t, streamName, false)
}
