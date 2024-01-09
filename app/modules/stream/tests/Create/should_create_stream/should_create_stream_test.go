package should_create_stream

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/golang_sse/app/ship/translator"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/AssertDoesStreamExists"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/AssertResponseMessage"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/GetMux"
	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

var mux = GetMux.GetMux()
var route = router.ByName(names.StreamCreate)
var streamName = "messages"
var expectedMessage = translator.Translate("success")

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

	tester.AssertSame(http.StatusOK, response.GetStatus())

	AssertResponseMessage.AssertResponseMessage(response, expectedMessage)

	AssertDoesStreamExists.AssertDoesStreamExists(t, streamName, true)
}
