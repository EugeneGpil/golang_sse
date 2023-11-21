package RequestAndAssertError

import "testing"

type RequestAndAssertErrorArgs struct {
	T *testing.T
	RouteName string
	RequestBody interface{}
	ExpectedErrorKeys []string
	ExpectedErrorCode int
}
