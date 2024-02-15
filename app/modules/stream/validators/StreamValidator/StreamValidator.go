package StreamValidator

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/consts/errorCodes"
	"github.com/EugeneGpil/golang_sse/app/ship/response"
	"github.com/EugeneGpil/golang_sse/app/ship/translator"

	requestPackage "github.com/EugeneGpil/request"
)

type StreamValidator struct {
	writer  http.ResponseWriter
	request *http.Request
}

func New(writer http.ResponseWriter, request *http.Request) StreamValidator {
	return StreamValidator{
		writer:  writer,
		request: request,
	}
}

func (streamValidator StreamValidator) Run() string {
	requestMethod := streamValidator.request.Method

	isQueryParams := requestMethod == http.MethodGet || requestMethod == http.MethodDelete

	if isQueryParams {
		return streamValidator.getStreamQueryParam()
	}

	return streamValidator.getStreamBodyParam()
}

func (streamValidator StreamValidator) getStreamQueryParam() string {
	queryParams := requestPackage.New(streamValidator.request).GetSimpleQueryParams()

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

	err := requestPackage.New(streamValidator.request).DecodeBody(&body)
	if err != nil {
		http.Error(streamValidator.writer, err.Error(), http.StatusBadRequest)

		return ""
	}

	if body.Stream == "" {
		streamValidator.writeError()

		return ""
	}

	return body.Stream
}

func (streamValidator StreamValidator) writeError() {
	response.WriteValidationErrors(streamValidator.writer, errorCodes.StreamNameIsRequired, map[string]string{
		"stream": translator.Translate("errors.empty_stream_is_not_allowed"),
	})
}
