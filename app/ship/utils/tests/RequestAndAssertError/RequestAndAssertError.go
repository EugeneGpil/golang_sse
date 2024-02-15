package RequestAndAssertError

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/GetMux"
	"github.com/EugeneGpil/router"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

func RequestAndAssertError(args RequestAndAssertErrorArgs) {
	tester.SetTester(args.T)

	mux := GetMux.GetMux()

	route := router.ByName(args.RouteName)

	response := httpTester.Request(httpTester.GetRequestDto{
		Method: route.Method,
		Url:    route.Url,
		Body:   args.RequestBody,
		Query:  args.Query,
	}, mux)

	tester.AssertSame(http.StatusUnprocessableEntity, response.GetStatus())

	var responseBody struct {
		Message string
		Errors  map[string]string
		Code    int
	}

	err := response.DecodeBody(&responseBody)

	tester.AssertNil(err)

	tester.AssertSame(responseBody.Message == "", false)

	tester.AssertSame(responseBody.Code, args.ExpectedErrorCode)

	for _, errorKey := range args.ExpectedErrorKeys {
		tester.AssertSame(responseBody.Errors[errorKey] == "", false)
	}
}
