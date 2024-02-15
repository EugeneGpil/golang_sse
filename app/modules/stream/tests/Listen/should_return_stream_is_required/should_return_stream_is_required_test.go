package should_return_stream_is_required_error

import (
	"testing"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/consts/errorCodes"
	"github.com/EugeneGpil/golang_sse/app/ship/router/names"
	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests/RequestAndAssertError"
)

func Test_should_return_stream_is_required_error(t *testing.T) {
	query := map[string]string{
		"stream": "",
	}

	expectedErrorKeys := []string{
		"stream",
	}

	RequestAndAssertError.RequestAndAssertError(RequestAndAssertError.RequestAndAssertErrorArgs{
		T:                 t,
		RouteName:         names.StreamRemove,
		Query:             query,
		ExpectedErrorKeys: expectedErrorKeys,
		ExpectedErrorCode: errorCodes.StreamNameIsRequired,
	})
}
