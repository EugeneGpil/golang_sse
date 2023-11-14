package controllers

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/modules/stream/consts/errorCodes"
	"github.com/EugeneGpil/golang_sse/app/ship/response"
	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
	"github.com/EugeneGpil/golang_sse/app/ship/translator"

	requestPackage "github.com/EugeneGpil/request"
)

func Remove(writer http.ResponseWriter, request *http.Request) {
	var body struct {
		Stream string
	}

	err := requestPackage.New(request).DecodeBody(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	if body.Stream == "" {
		response.WriteValidationErrors(writer, errorCodes.StreamNameIsRequired, map[string]string{
			"stream": translator.Translate("errors.empty_stream_is_not_allowed"),
		})

		return
	}

	sseServer.Get().RemoveStream(body.Stream)
}
