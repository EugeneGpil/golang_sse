package response

import (
	"net/http"

	"github.com/EugeneGpil/golang_sse/app/ship/translator"
	responsePackage "github.com/EugeneGpil/response"
)

func WriteSuccess(writer http.ResponseWriter) {
	response := responsePackage.New(writer)

	response.WriteSuccess(translator.Translate("Success"))
}
