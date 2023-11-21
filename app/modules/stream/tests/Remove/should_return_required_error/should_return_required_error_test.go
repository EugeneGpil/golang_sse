package should_return_required_error

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/consts/errorCodes"
	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/GetMux"
	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

var mux = GetMux.GetMux()
var route = router.ByName(names.StreamRemove)

func Test_should_return_required_error(t *testing.T) {
	tester.SetTester(t)

	response := httpTester.Request(httpTester.GetRequestDto{
		Method: route.Method,
		Url: route.Url,
		Body: struct {
			Stream string
		} {
			Stream: "",
		},
	}, mux)

	tester.AssertSame(http.StatusUnprocessableEntity, response.GetStatus())

	var responseBody struct {
		Message string
		Errors  map[string]string
		Code int
	}

	err := response.DecodeBody(&responseBody)

	tester.AssertNil(err)

	tester.AssertSame(responseBody.Message == "", false)
	tester.AssertSame(responseBody.Errors["stream"] == "", false)
	tester.AssertSame(responseBody.Code, errorCodes.StreamNameIsRequired)
}
