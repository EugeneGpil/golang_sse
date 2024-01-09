package AssertResponseMessage

import (
	"github.com/EugeneGpil/golang_sse/app/ship/interfaces/tests/TestResponse"
	"github.com/EugeneGpil/tester"
)

func AssertResponseMessage(response TestResponse.TestResponse, expectedMessage string) {
	responseBody := struct {
		Message string
	}{}

	err := response.DecodeBody(&responseBody)

	tester.AssertNil(err)

	tester.AssertSame(expectedMessage, responseBody.Message)
}
