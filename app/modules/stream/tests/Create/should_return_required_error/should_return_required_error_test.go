package should_return_required_error

import (
	"testing"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/consts/errorCodes"
	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/RequestAndAssertError"
)

func Test_should_return_required_error(t *testing.T) {
	requestBody := struct {
		Stream string
	}{
		Stream: "",
	}

	expectedErrorKeys := []string{
		"stream",
	}

	RequestAndAssertError.RequestAndAssertError(RequestAndAssertError.RequestAndAssertErrorArgs{
		T:                 t,
		RouteName:         names.StreamCreate,
		RequestBody:       requestBody,
		ExpectedErrorKeys: expectedErrorKeys,
		ExpectedErrorCode: errorCodes.StreamNameIsRequired,
	})
}
