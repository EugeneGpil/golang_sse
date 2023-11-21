package should_return_required_error

import (
	"testing"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/consts/errorCodes"
	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/GetMux"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/RequestAndAssertError"
	"github.com/EugeneGpil/router"
)

var mux = GetMux.GetMux()
var route = router.ByName(names.StreamCreate)

func Test_should_return_required_error(t *testing.T) {
	RequestAndAssertError.RequestAndAssertError(RequestAndAssertError.RequestAndAssertErrorArgs{
		T: t,
		RouteName: names.StreamCreate,
		RequestBody: struct {
			Stream string
		}{
			Stream: "",
		},
		ExpectedErrorKeys: []string{
			"stream",
		},
		ExpectedErrorCode: errorCodes.StreamNameIsRequired,
	})
}
