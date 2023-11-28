package AssertDoesStreamExists

import (
	"testing"

	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
	"github.com/EugeneGpil/tester"
)

func AssertDoesStreamExists(t *testing.T, streamName string) {
	tester.SetTester(t)

	doesStreamExistAfterRequest := sseServer.Get().StreamExists(streamName)

	tester.AssertSame(doesStreamExistAfterRequest, true)
}
