package StreamValidator

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/consts/errorCodes"
	"github.com/EugeneGpil/golang_sse/app/ship/response"
	"github.com/EugeneGpil/golang_sse/app/ship/translator"

	requestPackage "github.com/EugeneGpil/request"
	responsePackage "github.com/EugeneGpil/response"
)

type StreamValidator struct {
	response responsePackage.Response
	request  *requestPackage.Request
}

func New(response responsePackage.Response, request *requestPackage.Request) StreamValidator {
	return StreamValidator{
		response: response,
		request:  request,
	}
}

func (streamValidator StreamValidator) Run() string {
	requestMethod := streamValidator.request.GetMethod()

	isQueryParams := requestMethod == http.MethodGet || requestMethod == http.MethodDelete

	if isQueryParams {
		return streamValidator.getStreamQueryParam()
	}

	return streamValidator.getStreamBodyParam()
}

func (streamValidator StreamValidator) getStreamQueryParam() string {
	queryParams := streamValidator.request.GetSimpleQueryParams()

	stream, isset := queryParams["stream"]
	if isset == false || stream == "" {
		streamValidator.writeError()

		return ""
	}

	return stream
}

func (streamValidator StreamValidator) getStreamBodyParam() string {
	var body struct {
		Stream string
	}

	err := streamValidator.request.DecodeBody(&body)
	if err != nil {
		streamValidator.response.WriteBadRequest(err.Error())

		return ""
	}

	if body.Stream == "" {
		streamValidator.writeError()

		return ""
	}

	return body.Stream
}

func (streamValidator StreamValidator) writeError() {
	response.WriteValidationErrors(streamValidator.response, errorCodes.StreamNameIsRequired, map[string]string{
		"stream": translator.Translate("errors.empty_stream_is_not_allowed"),
	})
}
