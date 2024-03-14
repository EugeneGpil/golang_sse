package controllers

import (
	"github.com/EugeneGpil/golang_sse/app/modules/stream/validators/StreamValidator"
	"github.com/EugeneGpil/golang_sse/app/ship/sseServer"
	"github.com/EugeneGpil/request"

	shipResponse "github.com/EugeneGpil/golang_sse/app/ship/response"
	responsePackage "github.com/EugeneGpil/response"
)

func Create(response responsePackage.Response, request *request.Request) {
	stream := StreamValidator.New(response, request).Run()
	if stream == "" {
		return
	}

	sseServer.Get().CreateStream(stream)

	shipResponse.WriteSuccess(response)
}
