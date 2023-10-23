package should_return_required_validation_error

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/consts/errorCodes"
	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests"
	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

var mux = tests.GetMux()
var route = router.ByName(names.StreamCreate)

type ResponseBodyType struct {
	Message string
	Errors  map[string]string
	Code int
}

func Test_should_return_required_validation_error(t *testing.T) {
	tester.SetTester(t)

	response := httpTester.Request(httpTester.GetRequestDto{
		Method: route.Method,
		Url:    route.Url,
		Body: struct {
			Stream string
		}{
			Stream: "",
		},
	}, mux)

	tester.AssertSame(http.StatusUnprocessableEntity, response.GetStatus())

	responseBodyRaw := response.GetBody()

	var responseBody = ResponseBodyType{}

	err := json.Unmarshal(responseBodyRaw, &responseBody)

	tester.AssertNil(err)

	// tester.AssertFalse(responseBody.Message == "")
	// tester.AssertFalse(responseBody.Errors["stream"] == "")
	tester.AssertSame(responseBody.Code, errorCodes.StreamNameIsRequired)
}
