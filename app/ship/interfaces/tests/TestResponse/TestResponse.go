package TestResponse

type TestResponse interface{
	DecodeBody(body interface{}) error
}
