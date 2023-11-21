package should_return_required_error

import (
	"testing"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/consts/errorCodes"
	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/RequestAndAssertError"
)

func Test_should_return_required_error(t *testing.T) {
	RequestAndAssertError.RequestAndAssertError(RequestAndAssertError.RequestAndAssertErrorArgs{
		T: t,
		RouteName: names.StreamRemove,
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
