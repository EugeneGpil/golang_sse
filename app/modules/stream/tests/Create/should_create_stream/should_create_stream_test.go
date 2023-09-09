package should_create_stream

import (
	"testing"

	"github.com/EugeneGpil/golang_sse/app/ship/utils/tests"
	"github.com/EugeneGpil/tester"
)

var mux = tests.GetMux()

func Test_should_create_stream(t *testing.T) {
	tester.SetTester(t)

}
