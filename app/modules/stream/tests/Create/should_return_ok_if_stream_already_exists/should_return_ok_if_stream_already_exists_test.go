package should_return_ok_if_stream_already_exists

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
var route = router.ByName(names.StreamCreate)
var expectedMessage = translator.Translate("success")

func Test_should_return_ok_if_stream_already_exists(t *testing.T) {
	tester.SetTester(t)

	sseServer.Get().CreateStream(streamName)

	isStreamExists := sseServer.Get().StreamExists(streamName)

	tester.AssertSame(isStreamExists, true)

	body := struct{
		Stream string
	} {
		Stream: streamName,
	}

	response := httpTester.Request(httpTester.GetRequestDto{
		Method: route.Method,
		Url:    route.Url,
		Body:   body,
	}, mux)

	tester.AssertSame(http.StatusOK, response.GetStatus())

	AssertResponseMessage.AssertResponseMessage(response, expectedMessage)

	AssertDoesStreamExists.AssertDoesStreamExists(t, streamName, true)
}
